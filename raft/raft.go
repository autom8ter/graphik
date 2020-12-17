package raft

import (
	"errors"
	"fmt"
	"github.com/hashicorp/raft"
	"net"
	"os"
	"path/filepath"
)

type Raft struct {
	raft *raft.Raft
	opts *Options
}

func NewRaft(fsm *FSM, opts ...Opt) (*Raft, error) {
	if fsm == nil {
		return nil, errors.New("empty fsm")
	}
	options := &Options{}
	for _, o := range opts {
		o(options)
	}
	options.setDefaults()
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(options.peerID)
	addr, err := net.ResolveTCPAddr("tcp", options.listenAddr)
	if err != nil {
		return nil, err
	}
	transport, err := raft.NewTCPTransport(options.listenAddr, addr, options.maxPool, options.timeout, os.Stderr)
	if err != nil {
		return nil, err
	}

	// Create the snapshot store. This allows the Raft to truncate the log.
	snapshots, err := raft.NewFileSnapshotStore(options.raftDir, options.retainSnapshots, os.Stderr)
	if err != nil {
		return nil, err
	}
	strg, err := NewStorage(filepath.Join(options.raftDir, "raft.db"))
	if err != nil {
		return nil, err
	}
	ra, err := raft.NewRaft(config, fsm, strg, strg, snapshots, transport)
	if err != nil {
		return nil, err
	}
	return &Raft{
		opts: options,
		raft: ra,
	}, nil
}

func (s *Raft) Join(nodeID, addr string) error {
	configFuture := s.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		return err
	}

	for _, srv := range configFuture.Configuration().Servers {
		if srv.ID == raft.ServerID(nodeID) || srv.Address == raft.ServerAddress(addr) {
			// However if *both* the ID and the address are the same, then nothing -- not even
			// a join operation -- is needed.
			if srv.Address == raft.ServerAddress(addr) && srv.ID == raft.ServerID(nodeID) {
				// already a member
				return nil
			}

			future := s.raft.RemoveServer(srv.ID, 0, 0)
			if err := future.Error(); err != nil {
				return fmt.Errorf("error removing existing node %s at %s: %s", nodeID, addr, err)
			}
		}
	}

	f := s.raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
	if f.Error() != nil {
		return f.Error()
	}
	return nil
}

func (s *Raft) LeaderAddr() string {
	return string(s.raft.Leader())
}

func (s *Raft) Stats() map[string]string {
	return s.raft.Stats()
}
