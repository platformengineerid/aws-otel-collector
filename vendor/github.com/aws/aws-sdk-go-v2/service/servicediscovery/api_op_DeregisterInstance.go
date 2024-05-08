// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the Amazon Route 53 DNS records and health check, if any, that Cloud
// Map created for the specified instance.
func (c *Client) DeregisterInstance(ctx context.Context, params *DeregisterInstanceInput, optFns ...func(*Options)) (*DeregisterInstanceOutput, error) {
	if params == nil {
		params = &DeregisterInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterInstance", params, optFns, c.addOperationDeregisterInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterInstanceInput struct {

	// The value that you specified for Id in the RegisterInstance (https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html)
	// request.
	//
	// This member is required.
	InstanceId *string

	// The ID of the service that the instance is associated with.
	//
	// This member is required.
	ServiceId *string

	noSmithyDocumentSerde
}

type DeregisterInstanceOutput struct {

	// A value that you can use to determine whether the request completed
	// successfully. To get the status of the operation, see GetOperation (https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html)
	// .
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeregisterInstance"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeregisterInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterInstance(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeregisterInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeregisterInstance",
	}
}
