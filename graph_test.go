package graphik_test

import (
	"context"
	"github.com/autom8ter/graphik"
	"github.com/autom8ter/graphik/backends/boltdb"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func Test(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.RemoveAll(tmpdir)
	graph, err := graphik.New(boltdb.Open(tmpdir))
	if err != nil {
		t.Fatal(err.Error())
	}
	workerCounter := 0
	testWorker := graphik.NewWorker("testworker", func(g graphik.Graphik) error {
		friendsFromSchool, err := graphik.NewEdgeQuery().
			Mod(graphik.EdgeModHandler(func(g graphik.Graph, e graphik.Edge) error {
				workerCounter++
				t.Logf("worker edge(%v) = %s", workerCounter+1, e.String())
				return nil
			})).Validate()
		if err != nil {
			return err
		}
		if err := graph.QueryEdges(context.Background(), friendsFromSchool); err != nil {
			return err
		}
		return nil
	}, func(err error) {
		t.Fatal(err.Error())
	}, 1*time.Second)
	graph.AddWorkers(testWorker)
	graph.StartWorkers()
	defer graph.StopWorkers()
	coleman := graphik.NewNode(graphik.NewPath("user", "cword3"))
	coleman.SetAttribute("name", "coleman")
	if err := graph.AddNode(context.Background(), coleman); err != nil {
		t.Fatal(err.Error())
	}
	tyler := graphik.NewNode(graphik.NewPath("user", "tyler123"))
	tyler.SetAttribute("name", "tyler")
	if err := graph.AddNode(context.Background(), tyler); err != nil {
		t.Fatal(err.Error())
	}
	friendship := graphik.NewEdge(graphik.NewEdgePath(coleman, "friends", tyler))
	friendship.SetAttribute("source", "school")
	if err := graph.AddEdge(context.Background(), friendship); err != nil {
		t.Fatal(err.Error())
	}
	time.Sleep(5 * time.Second)
	if workerCounter == 0 {
		t.Fatal("worker didnt execute")
	}
}
