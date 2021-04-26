
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  name: oauth2clients.hydra.ory.sh
spec:
  group: hydra.ory.sh
  names:
    kind: OAuth2Client
    plural: oauth2clients
  scope: ""
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: OAuth2Client is the Schema for the oauth2clients API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        metadata:
          properties:
            annotations:
              additionalProperties:
                type: string
              description: 'Annotations is an unstructured key value map stored with
                a resource that may be set by external tools to store and retrieve
                arbitrary metadata. They are not queryable and should be preserved
                when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations'
              type: object
            clusterName:
              description: The name of the cluster which the object belongs to. This
                is used to distinguish resources with same name and namespace in different
                clusters. This field is not set anywhere right now and apiserver is
                going to ignore it if set in create or update request.
              type: string
            creationTimestamp:
              description: "CreationTimestamp is a timestamp representing the server
                time when this object was created. It is not guaranteed to be set
                in happens-before order across separate operations. Clients may not
                set this value. It is represented in RFC3339 form and is in UTC. \n
                Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
              format: date-time
              type: string
            deletionGracePeriodSeconds:
              description: Number of seconds allowed for this object to gracefully
                terminate before it will be removed from the system. Only set when
                deletionTimestamp is also set. May only be shortened. Read-only.
              format: int64
              type: integer
            deletionTimestamp:
              description: "DeletionTimestamp is RFC 3339 date and time at which this
                resource will be deleted. This field is set by the server when a graceful
                deletion is requested by the user, and is not directly settable by
                a client. The resource is expected to be deleted (no longer visible
                from resource lists, and not reachable by name) after the time in
                this field, once the finalizers list is empty. As long as the finalizers
                list contains items, deletion is blocked. Once the deletionTimestamp
                is set, this value may not be unset or be set further into the future,
                although it may be shortened or the resource may be deleted prior
                to this time. For example, a user may request that a pod is deleted
                in 30 seconds. The Kubelet will react by sending a graceful termination
                signal to the containers in the pod. After that 30 seconds, the Kubelet
                will send a hard termination signal (SIGKILL) to the container and
                after cleanup, remove the pod from the API. In the presence of network
                partitions, this object may still exist after this timestamp, until
                an administrator or automated process can determine the resource is
                fully terminated. If not set, graceful deletion of the object has
                not been requested. \n Populated by the system when a graceful deletion
                is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
              format: date-time
              type: string
            finalizers:
              description: Must be empty before the object is deleted from the registry.
                Each entry is an identifier for the responsible component that will
                remove the entry from the list. If the deletionTimestamp of the object
                is non-nil, entries in this list can only be removed.
              items:
                type: string
              type: array
            generateName:
              description: "GenerateName is an optional prefix, used by the server,
                to generate a unique name ONLY IF the Name field has not been provided.
                If this field is used, the name returned to the client will be different
                than the name passed. This value will also be combined with a unique
                suffix. The provided value has the same validation rules as the Name
                field, and may be truncated by the length of the suffix required to
                make the value unique on the server. \n If this field is specified
                and the generated name exists, the server will NOT return a 409 -
                instead, it will either return 201 Created or 500 with Reason ServerTimeout
                indicating a unique name could not be found in the time allotted,
                and the client should retry (optionally after the time indicated in
                the Retry-After header). \n Applied only if Name is not specified.
                More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency"
              type: string
            generation:
              description: A sequence number representing a specific generation of
                the desired state. Populated by the system. Read-only.
              format: int64
              type: integer
            initializers:
              description: "An initializer is a controller which enforces some system
                invariant at object creation time. This field is a list of initializers
                that have not yet acted on this object. If nil or empty, this object
                has been completely initialized. Otherwise, the object is considered
                uninitialized and is hidden (in list/watch and get calls) from clients
                that haven't explicitly asked to observe uninitialized objects. \n
                When an object is created, the system will populate this list with
                the current set of initializers. Only privileged users may set or
                modify this list. Once it is empty, it may not be modified further
                by any user. \n DEPRECATED - initializers are an alpha field and will
                be removed in v1.15."
              properties:
                pending:
                  description: Pending is a list of initializers that must execute
                    in order before this object is visible. When the last pending
                    initializer is removed, and no failing result is set, the initializers
                    struct will be set to nil and the object is considered as initialized
                    and visible to all clients.
                  items:
                    properties:
                      name:
                        description: name of the process that is responsible for initializing
                          this object.
                        type: string
                    required:
                    - name
                    type: object
                  type: array
                result:
                  description: If result is set with the Failure field, the object
                    will be persisted to storage and then deleted, ensuring that other
                    clients can observe the deletion.
                  properties:
                    apiVersion:
                      description: 'APIVersion defines the versioned schema of this
                        representation of an object. Servers should convert recognized
                        schemas to the latest internal value, and may reject unrecognized
                        values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
                      type: string
                    code:
                      description: Suggested HTTP return code for this status, 0 if
                        not set.
                      format: int32
                      type: integer
                    details:
                      description: Extended data associated with the reason.  Each
                        reason may define its own extended details. This field is
                        optional and the data returned is not guaranteed to conform
                        to any schema except that defined by the reason type.
                      properties:
                        causes:
                          description: The Causes array includes more details associated
                            with the StatusReason failure. Not all StatusReasons may
                            provide detailed causes.
                          items:
                            properties:
                              field:
                                description: "The field of the resource that has caused
                                  this error, as named by its JSON serialization.
                                  May include dot and postfix notation for nested
                                  attributes. Arrays are zero-indexed.  Fields may
                                  appear more than once in an array of causes due
                                  to fields having multiple errors. Optional. \n Examples:
                                  \  \"name\" - the field \"name\" on the current
                                  resource   \"items[0].name\" - the field \"name\"
                                  on the first array entry in \"items\""
                                type: string
                              message:
                                description: A human-readable description of the cause
                                  of the error.  This field may be presented as-is
                                  to a reader.
                                type: string
                              reason:
                                description: A machine-readable description of the
                                  cause of the error. If this value is empty there
                                  is no information available.
                                type: string
                            type: object
                          type: array
                        group:
                          description: The group attribute of the resource associated
                            with the status StatusReason.
                          type: string
                        kind:
                          description: 'The kind attribute of the resource associated
                            with the status StatusReason. On some operations may differ
                            from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: The name attribute of the resource associated
                            with the status StatusReason (when there is a single name
                            which can be described).
                          type: string
                        retryAfterSeconds:
                          description: If specified, the time in seconds before the
                            operation should be retried. Some errors may indicate
                            the client must take an alternate action - for those errors
                            this field may indicate how long to wait before taking
                            the alternate action.
                          format: int32
                          type: integer
                        uid:
                          description: 'UID of the resource. (when there is a single
                            resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                          type: string
                      type: object
                    kind:
                      description: 'Kind is a string value representing the REST resource
                        this object represents. Servers may infer this from the endpoint
                        the client submits requests to. Cannot be updated. In CamelCase.
                        More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                      type: string
                    message:
                      description: A human-readable description of the status of this
                        operation.
                      type: string
                    metadata:
                      description: 'Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                      properties:
                        continue:
                          description: continue may be set if the user set a limit
                            on the number of items returned, and indicates that the
                            server has more data available. The value is opaque and
                            may be used to issue another request to the endpoint that
                            served this list to retrieve the next set of available
                            objects. Continuing a consistent list may not be possible
                            if the server configuration has changed or more than a
                            few minutes have passed. The resourceVersion field returned
                            when using this continue value will be identical to the
                            value in the first response, unless you have received
                            this token from an error message.
                          type: string
                        resourceVersion:
                          description: 'String that identifies the server''s internal
                            version of this object that can be used by clients to
                            determine when objects have changed. Value must be treated
                            as opaque by clients and passed unmodified back to the
                            server. Populated by the system. Read-only. More info:
                            https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        selfLink:
                          description: selfLink is a URL representing this object.
                            Populated by the system. Read-only.
                          type: string
                      type: object
                    reason:
                      description: A machine-readable description of why this operation
                        is in the "Failure" status. If this value is empty there is
                        no information available. A Reason clarifies an HTTP status
                        code but does not override it.
                      type: string
                    status:
                      description: 'Status of the operation. One of: "Success" or
                        "Failure". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status'
                      type: string
                  type: object
              required:
              - pending
              type: object
            labels:
              additionalProperties:
                type: string
              description: 'Map of string keys and values that can be used to organize
                and categorize (scope and select) objects. May match selectors of
                replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels'
              type: object
            managedFields:
              description: "ManagedFields maps workflow-id and version to the set
                of fields that are managed by that workflow. This is mostly for internal
                housekeeping, and users typically shouldn't need to set or understand
                this field. A workflow can be the user's name, a controller's name,
                or the name of a specific apply path like \"ci-cd\". The set of fields
                is always in the version that the workflow used when modifying the
                object. \n This field is alpha and can be changed or removed without
                notice."
              items:
                properties:
                  apiVersion:
                    description: APIVersion defines the version of this resource that
                      this field set applies to. The format is "group/version" just
                      like the top-level APIVersion field. It is necessary to track
                      the version of a field set because it cannot be automatically
                      converted.
                    type: string
                  fields:
                    additionalProperties: true
                    description: Fields identifies a set of fields.
                    type: object
                  manager:
                    description: Manager is an identifier of the workflow managing
                      these fields.
                    type: string
                  operation:
                    description: Operation is the type of operation which lead to
                      this ManagedFieldsEntry being created. The only valid values
                      for this field are 'Apply' and 'Update'.
                    type: string
                  time:
                    description: Time is timestamp of when these fields were set.
                      It should always be empty if Operation is 'Apply'
                    format: date-time
                    type: string
                type: object
              type: array
            name:
              description: 'Name must be unique within a namespace. Is required when
                creating resources, although some resources may allow a client to
                request the generation of an appropriate name automatically. Name
                is primarily intended for creation idempotence and configuration definition.
                Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
              type: string
            namespace:
              description: "Namespace defines the space within each name must be unique.
                An empty namespace is equivalent to the \"default\" namespace, but
                \"default\" is the canonical representation. Not all objects are required
                to be scoped to a namespace - the value of this field for those objects
                will be empty. \n Must be a DNS_LABEL. Cannot be updated. More info:
                http://kubernetes.io/docs/user-guide/namespaces"
              type: string
            ownerReferences:
              description: List of objects depended by this object. If ALL objects
                in the list have been deleted, this object will be garbage collected.
                If this object is managed by a controller, then an entry in this list
                will point to this controller, with the controller field set to true.
                There cannot be more than one managing controller.
              items:
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  blockOwnerDeletion:
                    description: If true, AND if the owner has the "foregroundDeletion"
                      finalizer, then the owner cannot be deleted from the key-value
                      store until this reference is removed. Defaults to false. To
                      set this field, a user needs "delete" permission of the owner,
                      otherwise 422 (Unprocessable Entity) will be returned.
                    type: boolean
                  controller:
                    description: If true, this reference points to the managing controller.
                    type: boolean
                  kind:
                    description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                    type: string
                  name:
                    description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                    type: string
                  uid:
                    description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                    type: string
                required:
                - apiVersion
                - kind
                - name
                - uid
                type: object
              type: array
            resourceVersion:
              description: "An opaque value that represents the internal version of
                this object that can be used by clients to determine when objects
                have changed. May be used for optimistic concurrency, change detection,
                and the watch operation on a resource or set of resources. Clients
                must treat these values as opaque and passed unmodified back to the
                server. They may only be valid for a particular resource or set of
                resources. \n Populated by the system. Read-only. Value must be treated
                as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency"
              type: string
            selfLink:
              description: SelfLink is a URL representing this object. Populated by
                the system. Read-only.
              type: string
            uid:
              description: "UID is the unique in time and space value for this object.
                It is typically generated by the server on successful creation of
                a resource and is not allowed to change on PUT operations. \n Populated
                by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids"
              type: string
          type: object
        spec:
          properties:
            allowedCorsOrigins:
              description: AllowedCorsOrigins is an array of allowed CORS origins
              items:
                pattern: \w+:/?/?[^\s]+
                type: string
              type: array
            audience:
              description: Audience is a whitelist defining the audiences this client
                is allowed to request tokens for
              items:
                type: string
              type: array
            clientName:
              description: ClientName is the human-readable string name of the client
                to be presented to the end-user during authorization.
              type: string
            grantTypes:
              description: GrantTypes is an array of grant types the client is allowed
                to use.
              items:
                enum:
                - client_credentials
                - authorization_code
                - implicit
                - refresh_token
                type: string
              maxItems: 4
              minItems: 1
              type: array
            hydraAdmin:
              description: HydraAdmin is the optional configuration to use for managing
                this client
              properties:
                endpoint:
                  description: Endpoint is the endpoint for the hydra instance on
                    which to set up the client. This value will override the value
                    provided to `--endpoint` (defaults to `"/clients"` in the application)
                  pattern: (^$|^/.*)
                  type: string
                forwardedProto:
                  description: ForwardedProto overrides the `--forwarded-proto` flag.
                    The value "off" will force this to be off even if `--forwarded-proto`
                    is specified
                  pattern: (^$|https?|off)
                  type: string
                port:
                  description: Port is the port for the hydra instance on which to
                    set up the client. This value will override the value provided
                    to `--hydra-port`
                  maximum: 65535
                  type: integer
                url:
                  description: URL is the URL for the hydra instance on which to set
                    up the client. This value will override the value provided to
                    `--hydra-url`
                  maxLength: 64
                  pattern: (^$|^https?://.*)
                  type: string
              type: object
            metadata:
              description: Metadata is abritrary data
              format: byte
              type: string
            postLogoutRedirectUris:
              description: PostLogoutRedirectURIs is an array of the post logout redirect
                URIs allowed for the application
              items:
                pattern: \w+:/?/?[^\s]+
                type: string
              type: array
            redirectUris:
              description: RedirectURIs is an array of the redirect URIs allowed for
                the application
              items:
                pattern: \w+:/?/?[^\s]+
                type: string
              type: array
            responseTypes:
              description: ResponseTypes is an array of the OAuth 2.0 response type
                strings that the client can use at the authorization endpoint.
              items:
                enum:
                - id_token
                - code
                - token
                type: string
              maxItems: 3
              minItems: 1
              type: array
            scope:
              description: Scope is a string containing a space-separated list of
                scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749])
                that the client can use when requesting access tokens.
              pattern: ([a-zA-Z0-9\.\*]+\s?)+
              type: string
            secretName:
              description: SecretName points to the K8s secret that contains this
                client's ID and password
              maxLength: 253
              minLength: 1
              pattern: '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*'
              type: string
            tokenEndpointAuthMethod:
              description: Indication which authentication method shoud be used for
                the token endpoint
              enum:
              - client_secret_basic
              - client_secret_post
              - private_key_jwt
              - none
              type: string
            frontchannelLogoutURI:
              description: Client URL that will cause the client to log itself out when
                rendered in an iframe by Hydra.
              type: string
            frontchannelLogoutSessionRequired:
              description: Boolean flag specifying whether the client requires that a 
                sid (session ID) Claim be included in the Logout Token to identify the 
                client session with the OP when the frontchannel-logout-callback is used. 
                If omitted, the default value is false.
              type: boolean
            backchannelLogoutURI:
              description: Client URL that will cause the client to log itself out when
                sent a Logout Token by Hydra.
              type: string
            backchannelLogoutSessionRequired:
              description: Boolean flag specifying whether the client requires that a
                sid (session ID) Claim be included in the Logout Token to identify the
                client session with the OP when the backchannel-logout-callback is used.
                If omitted, the default value is false.
              type: boolean
          required:
          - grantTypes
          - scope
          - secretName
          type: object
        status:
          properties:
            observedGeneration:
              description: ObservedGeneration represents the most recent generation
                observed by the daemon set controller.
              format: int64
              type: integer
            reconciliationError:
              properties:
                description:
                  description: Description is the description of the reconciliation
                    error
                  type: string
                statusCode:
                  description: Code is the status code of the reconciliation error
                  type: string
              type: object
          type: object
      type: object
  versions:
  - name: v1alpha1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
