<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: graphik.proto

namespace Api;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Edit patches the attributes of a Doc or Connection
 *
 * Generated from protobuf message <code>api.Edit</code>
 */
class Edit extends \Google\Protobuf\Internal\Message
{
    /**
     * ref is the ref to the target doc/connection to patch
     *
     * Generated from protobuf field <code>.api.Ref ref = 1 [(.validator.field) = {</code>
     */
    private $ref = null;
    /**
     * attributes are k/v pairs used to overwrite k/v pairs on a doc/connection
     *
     * Generated from protobuf field <code>.google.protobuf.Struct attributes = 2;</code>
     */
    private $attributes = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Api\Ref $ref
     *           ref is the ref to the target doc/connection to patch
     *     @type \Google\Protobuf\Struct $attributes
     *           attributes are k/v pairs used to overwrite k/v pairs on a doc/connection
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Graphik::initOnce();
        parent::__construct($data);
    }

    /**
     * ref is the ref to the target doc/connection to patch
     *
     * Generated from protobuf field <code>.api.Ref ref = 1 [(.validator.field) = {</code>
     * @return \Api\Ref
     */
    public function getRef()
    {
        return $this->ref;
    }

    /**
     * ref is the ref to the target doc/connection to patch
     *
     * Generated from protobuf field <code>.api.Ref ref = 1 [(.validator.field) = {</code>
     * @param \Api\Ref $var
     * @return $this
     */
    public function setRef($var)
    {
        GPBUtil::checkMessage($var, \Api\Ref::class);
        $this->ref = $var;

        return $this;
    }

    /**
     * attributes are k/v pairs used to overwrite k/v pairs on a doc/connection
     *
     * Generated from protobuf field <code>.google.protobuf.Struct attributes = 2;</code>
     * @return \Google\Protobuf\Struct
     */
    public function getAttributes()
    {
        return $this->attributes;
    }

    /**
     * attributes are k/v pairs used to overwrite k/v pairs on a doc/connection
     *
     * Generated from protobuf field <code>.google.protobuf.Struct attributes = 2;</code>
     * @param \Google\Protobuf\Struct $var
     * @return $this
     */
    public function setAttributes($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Struct::class);
        $this->attributes = $var;

        return $this;
    }

}

