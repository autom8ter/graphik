<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: graphik.proto

namespace Api;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>api.RaftCommand</code>
 */
class RaftCommand extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.api.Doc user = 1;</code>
     */
    private $user = null;
    /**
     * Generated from protobuf field <code>string method = 2;</code>
     */
    private $method = '';
    /**
     * Generated from protobuf field <code>repeated .api.Doc set_docs = 3;</code>
     */
    private $set_docs;
    /**
     * Generated from protobuf field <code>repeated .api.Connection set_connections = 4;</code>
     */
    private $set_connections;
    /**
     * Generated from protobuf field <code>repeated .api.Ref del_docs = 5;</code>
     */
    private $del_docs;
    /**
     * Generated from protobuf field <code>repeated .api.Ref del_connections = 6;</code>
     */
    private $del_connections;
    /**
     * Generated from protobuf field <code>.api.Indexes set_indexes = 7;</code>
     */
    private $set_indexes = null;
    /**
     * Generated from protobuf field <code>.api.Authorizers set_authorizers = 8;</code>
     */
    private $set_authorizers = null;
    /**
     * Generated from protobuf field <code>.api.TypeValidators set_type_validators = 9;</code>
     */
    private $set_type_validators = null;
    /**
     * Generated from protobuf field <code>.api.Message send_message = 10;</code>
     */
    private $send_message = null;
    /**
     * Generated from protobuf field <code>.api.Triggers set_triggers = 11;</code>
     */
    private $set_triggers = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Api\Doc $user
     *     @type string $method
     *     @type \Api\Doc[]|\Google\Protobuf\Internal\RepeatedField $set_docs
     *     @type \Api\Connection[]|\Google\Protobuf\Internal\RepeatedField $set_connections
     *     @type \Api\Ref[]|\Google\Protobuf\Internal\RepeatedField $del_docs
     *     @type \Api\Ref[]|\Google\Protobuf\Internal\RepeatedField $del_connections
     *     @type \Api\Indexes $set_indexes
     *     @type \Api\Authorizers $set_authorizers
     *     @type \Api\TypeValidators $set_type_validators
     *     @type \Api\Message $send_message
     *     @type \Api\Triggers $set_triggers
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Graphik::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.api.Doc user = 1;</code>
     * @return \Api\Doc
     */
    public function getUser()
    {
        return $this->user;
    }

    /**
     * Generated from protobuf field <code>.api.Doc user = 1;</code>
     * @param \Api\Doc $var
     * @return $this
     */
    public function setUser($var)
    {
        GPBUtil::checkMessage($var, \Api\Doc::class);
        $this->user = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string method = 2;</code>
     * @return string
     */
    public function getMethod()
    {
        return $this->method;
    }

    /**
     * Generated from protobuf field <code>string method = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setMethod($var)
    {
        GPBUtil::checkString($var, True);
        $this->method = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .api.Doc set_docs = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getSetDocs()
    {
        return $this->set_docs;
    }

    /**
     * Generated from protobuf field <code>repeated .api.Doc set_docs = 3;</code>
     * @param \Api\Doc[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setSetDocs($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Api\Doc::class);
        $this->set_docs = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .api.Connection set_connections = 4;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getSetConnections()
    {
        return $this->set_connections;
    }

    /**
     * Generated from protobuf field <code>repeated .api.Connection set_connections = 4;</code>
     * @param \Api\Connection[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setSetConnections($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Api\Connection::class);
        $this->set_connections = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .api.Ref del_docs = 5;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getDelDocs()
    {
        return $this->del_docs;
    }

    /**
     * Generated from protobuf field <code>repeated .api.Ref del_docs = 5;</code>
     * @param \Api\Ref[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setDelDocs($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Api\Ref::class);
        $this->del_docs = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .api.Ref del_connections = 6;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getDelConnections()
    {
        return $this->del_connections;
    }

    /**
     * Generated from protobuf field <code>repeated .api.Ref del_connections = 6;</code>
     * @param \Api\Ref[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setDelConnections($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Api\Ref::class);
        $this->del_connections = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.api.Indexes set_indexes = 7;</code>
     * @return \Api\Indexes
     */
    public function getSetIndexes()
    {
        return $this->set_indexes;
    }

    /**
     * Generated from protobuf field <code>.api.Indexes set_indexes = 7;</code>
     * @param \Api\Indexes $var
     * @return $this
     */
    public function setSetIndexes($var)
    {
        GPBUtil::checkMessage($var, \Api\Indexes::class);
        $this->set_indexes = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.api.Authorizers set_authorizers = 8;</code>
     * @return \Api\Authorizers
     */
    public function getSetAuthorizers()
    {
        return $this->set_authorizers;
    }

    /**
     * Generated from protobuf field <code>.api.Authorizers set_authorizers = 8;</code>
     * @param \Api\Authorizers $var
     * @return $this
     */
    public function setSetAuthorizers($var)
    {
        GPBUtil::checkMessage($var, \Api\Authorizers::class);
        $this->set_authorizers = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.api.TypeValidators set_type_validators = 9;</code>
     * @return \Api\TypeValidators
     */
    public function getSetTypeValidators()
    {
        return $this->set_type_validators;
    }

    /**
     * Generated from protobuf field <code>.api.TypeValidators set_type_validators = 9;</code>
     * @param \Api\TypeValidators $var
     * @return $this
     */
    public function setSetTypeValidators($var)
    {
        GPBUtil::checkMessage($var, \Api\TypeValidators::class);
        $this->set_type_validators = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.api.Message send_message = 10;</code>
     * @return \Api\Message
     */
    public function getSendMessage()
    {
        return $this->send_message;
    }

    /**
     * Generated from protobuf field <code>.api.Message send_message = 10;</code>
     * @param \Api\Message $var
     * @return $this
     */
    public function setSendMessage($var)
    {
        GPBUtil::checkMessage($var, \Api\Message::class);
        $this->send_message = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.api.Triggers set_triggers = 11;</code>
     * @return \Api\Triggers
     */
    public function getSetTriggers()
    {
        return $this->set_triggers;
    }

    /**
     * Generated from protobuf field <code>.api.Triggers set_triggers = 11;</code>
     * @param \Api\Triggers $var
     * @return $this
     */
    public function setSetTriggers($var)
    {
        GPBUtil::checkMessage($var, \Api\Triggers::class);
        $this->set_triggers = $var;

        return $this;
    }

}
