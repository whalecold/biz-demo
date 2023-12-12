// Code generated by thriftgo (0.2.9). DO NOT EDIT.

package details

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type Product struct {
	ID          string `thrift:"ID,1,required" frugal:"1,required,string" json:"ID"`
	Title       string `thrift:"Title,2,required" frugal:"2,required,string" json:"Title"`
	Author      string `thrift:"Author,3,required" frugal:"3,required,string" json:"Author"`
	Description string `thrift:"Description,4,required" frugal:"4,required,string" json:"Description"`
	AuthorLink  string `thrift:"AuthorLink,5,required" frugal:"5,required,string" json:"AuthorLink"`
	Link        string `thrift:"Link,6,required" frugal:"6,required,string" json:"Link"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) InitDefault() {
	*p = Product{}
}

func (p *Product) GetID() (v string) {
	return p.ID
}

func (p *Product) GetTitle() (v string) {
	return p.Title
}

func (p *Product) GetAuthor() (v string) {
	return p.Author
}

func (p *Product) GetDescription() (v string) {
	return p.Description
}

func (p *Product) GetAuthorLink() (v string) {
	return p.AuthorLink
}

func (p *Product) GetLink() (v string) {
	return p.Link
}
func (p *Product) SetID(val string) {
	p.ID = val
}
func (p *Product) SetTitle(val string) {
	p.Title = val
}
func (p *Product) SetAuthor(val string) {
	p.Author = val
}
func (p *Product) SetDescription(val string) {
	p.Description = val
}
func (p *Product) SetAuthorLink(val string) {
	p.AuthorLink = val
}
func (p *Product) SetLink(val string) {
	p.Link = val
}

var fieldIDToName_Product = map[int16]string{
	1: "ID",
	2: "Title",
	3: "Author",
	4: "Description",
	5: "AuthorLink",
	6: "Link",
}

func (p *Product) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetID bool = false
	var issetTitle bool = false
	var issetAuthor bool = false
	var issetDescription bool = false
	var issetAuthorLink bool = false
	var issetLink bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetID = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetTitle = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
				issetAuthor = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
				issetDescription = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
				issetAuthorLink = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField6(iprot); err != nil {
					goto ReadFieldError
				}
				issetLink = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetID {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetTitle {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetAuthor {
		fieldId = 3
		goto RequiredFieldNotSetError
	}

	if !issetDescription {
		fieldId = 4
		goto RequiredFieldNotSetError
	}

	if !issetAuthorLink {
		fieldId = 5
		goto RequiredFieldNotSetError
	}

	if !issetLink {
		fieldId = 6
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Product[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_Product[fieldId]))
}

func (p *Product) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.ID = v
	}
	return nil
}

func (p *Product) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Title = v
	}
	return nil
}

func (p *Product) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Author = v
	}
	return nil
}

func (p *Product) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Description = v
	}
	return nil
}

func (p *Product) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.AuthorLink = v
	}
	return nil
}

func (p *Product) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Link = v
	}
	return nil
}

func (p *Product) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Product"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}
		if err = p.writeField6(oprot); err != nil {
			fieldId = 6
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Product) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("ID", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.ID); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Product) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Title", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Title); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Product) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Author", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Author); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Product) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Description", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Description); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *Product) writeField5(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("AuthorLink", thrift.STRING, 5); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.AuthorLink); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *Product) writeField6(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Link", thrift.STRING, 6); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Link); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 end error: ", p), err)
}

func (p *Product) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Product(%+v)", *p)
}

func (p *Product) DeepEqual(ano *Product) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.ID) {
		return false
	}
	if !p.Field2DeepEqual(ano.Title) {
		return false
	}
	if !p.Field3DeepEqual(ano.Author) {
		return false
	}
	if !p.Field4DeepEqual(ano.Description) {
		return false
	}
	if !p.Field5DeepEqual(ano.AuthorLink) {
		return false
	}
	if !p.Field6DeepEqual(ano.Link) {
		return false
	}
	return true
}

func (p *Product) Field1DeepEqual(src string) bool {

	if strings.Compare(p.ID, src) != 0 {
		return false
	}
	return true
}
func (p *Product) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Title, src) != 0 {
		return false
	}
	return true
}
func (p *Product) Field3DeepEqual(src string) bool {

	if strings.Compare(p.Author, src) != 0 {
		return false
	}
	return true
}
func (p *Product) Field4DeepEqual(src string) bool {

	if strings.Compare(p.Description, src) != 0 {
		return false
	}
	return true
}
func (p *Product) Field5DeepEqual(src string) bool {

	if strings.Compare(p.AuthorLink, src) != 0 {
		return false
	}
	return true
}
func (p *Product) Field6DeepEqual(src string) bool {

	if strings.Compare(p.Link, src) != 0 {
		return false
	}
	return true
}

type GetProductReq struct {
	ID string `thrift:"ID,1,required" frugal:"1,required,string" json:"ID"`
}

func NewGetProductReq() *GetProductReq {
	return &GetProductReq{}
}

func (p *GetProductReq) InitDefault() {
	*p = GetProductReq{}
}

func (p *GetProductReq) GetID() (v string) {
	return p.ID
}
func (p *GetProductReq) SetID(val string) {
	p.ID = val
}

var fieldIDToName_GetProductReq = map[int16]string{
	1: "ID",
}

func (p *GetProductReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetID bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetID = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetID {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetProductReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GetProductReq[fieldId]))
}

func (p *GetProductReq) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.ID = v
	}
	return nil
}

func (p *GetProductReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetProductReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetProductReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("ID", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.ID); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetProductReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetProductReq(%+v)", *p)
}

func (p *GetProductReq) DeepEqual(ano *GetProductReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.ID) {
		return false
	}
	return true
}

func (p *GetProductReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.ID, src) != 0 {
		return false
	}
	return true
}

type GetProductResp struct {
	Product *Product `thrift:"Product,1,required" frugal:"1,required,Product" json:"Product"`
}

func NewGetProductResp() *GetProductResp {
	return &GetProductResp{}
}

func (p *GetProductResp) InitDefault() {
	*p = GetProductResp{}
}

var GetProductResp_Product_DEFAULT *Product

func (p *GetProductResp) GetProduct() (v *Product) {
	if !p.IsSetProduct() {
		return GetProductResp_Product_DEFAULT
	}
	return p.Product
}
func (p *GetProductResp) SetProduct(val *Product) {
	p.Product = val
}

var fieldIDToName_GetProductResp = map[int16]string{
	1: "Product",
}

func (p *GetProductResp) IsSetProduct() bool {
	return p.Product != nil
}

func (p *GetProductResp) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetProduct bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetProduct = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetProduct {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetProductResp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GetProductResp[fieldId]))
}

func (p *GetProductResp) ReadField1(iprot thrift.TProtocol) error {
	p.Product = NewProduct()
	if err := p.Product.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *GetProductResp) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetProductResp"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetProductResp) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Product", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Product.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetProductResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetProductResp(%+v)", *p)
}

func (p *GetProductResp) DeepEqual(ano *GetProductResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Product) {
		return false
	}
	return true
}

func (p *GetProductResp) Field1DeepEqual(src *Product) bool {

	if !p.Product.DeepEqual(src) {
		return false
	}
	return true
}

type DetailsService interface {
	GetProduct(ctx context.Context, req *GetProductReq) (r *GetProductResp, err error)
}

type DetailsServiceClient struct {
	c thrift.TClient
}

func NewDetailsServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *DetailsServiceClient {
	return &DetailsServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewDetailsServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *DetailsServiceClient {
	return &DetailsServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewDetailsServiceClient(c thrift.TClient) *DetailsServiceClient {
	return &DetailsServiceClient{
		c: c,
	}
}

func (p *DetailsServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *DetailsServiceClient) GetProduct(ctx context.Context, req *GetProductReq) (r *GetProductResp, err error) {
	var _args DetailsServiceGetProductArgs
	_args.Req = req
	var _result DetailsServiceGetProductResult
	if err = p.Client_().Call(ctx, "GetProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type DetailsServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      DetailsService
}

func (p *DetailsServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *DetailsServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *DetailsServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewDetailsServiceProcessor(handler DetailsService) *DetailsServiceProcessor {
	self := &DetailsServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("GetProduct", &detailsServiceProcessorGetProduct{handler: handler})
	return self
}
func (p *DetailsServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type detailsServiceProcessorGetProduct struct {
	handler DetailsService
}

func (p *detailsServiceProcessorGetProduct) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := DetailsServiceGetProductArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetProduct", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := DetailsServiceGetProductResult{}
	var retval *GetProductResp
	if retval, err2 = p.handler.GetProduct(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetProduct: "+err2.Error())
		oprot.WriteMessageBegin("GetProduct", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetProduct", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type DetailsServiceGetProductArgs struct {
	Req *GetProductReq `thrift:"req,1" frugal:"1,default,GetProductReq" json:"req"`
}

func NewDetailsServiceGetProductArgs() *DetailsServiceGetProductArgs {
	return &DetailsServiceGetProductArgs{}
}

func (p *DetailsServiceGetProductArgs) InitDefault() {
	*p = DetailsServiceGetProductArgs{}
}

var DetailsServiceGetProductArgs_Req_DEFAULT *GetProductReq

func (p *DetailsServiceGetProductArgs) GetReq() (v *GetProductReq) {
	if !p.IsSetReq() {
		return DetailsServiceGetProductArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *DetailsServiceGetProductArgs) SetReq(val *GetProductReq) {
	p.Req = val
}

var fieldIDToName_DetailsServiceGetProductArgs = map[int16]string{
	1: "req",
}

func (p *DetailsServiceGetProductArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DetailsServiceGetProductArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DetailsServiceGetProductArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *DetailsServiceGetProductArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewGetProductReq()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *DetailsServiceGetProductArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetProduct_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *DetailsServiceGetProductArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *DetailsServiceGetProductArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DetailsServiceGetProductArgs(%+v)", *p)
}

func (p *DetailsServiceGetProductArgs) DeepEqual(ano *DetailsServiceGetProductArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *DetailsServiceGetProductArgs) Field1DeepEqual(src *GetProductReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type DetailsServiceGetProductResult struct {
	Success *GetProductResp `thrift:"success,0,optional" frugal:"0,optional,GetProductResp" json:"success,omitempty"`
}

func NewDetailsServiceGetProductResult() *DetailsServiceGetProductResult {
	return &DetailsServiceGetProductResult{}
}

func (p *DetailsServiceGetProductResult) InitDefault() {
	*p = DetailsServiceGetProductResult{}
}

var DetailsServiceGetProductResult_Success_DEFAULT *GetProductResp

func (p *DetailsServiceGetProductResult) GetSuccess() (v *GetProductResp) {
	if !p.IsSetSuccess() {
		return DetailsServiceGetProductResult_Success_DEFAULT
	}
	return p.Success
}
func (p *DetailsServiceGetProductResult) SetSuccess(x interface{}) {
	p.Success = x.(*GetProductResp)
}

var fieldIDToName_DetailsServiceGetProductResult = map[int16]string{
	0: "success",
}

func (p *DetailsServiceGetProductResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DetailsServiceGetProductResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DetailsServiceGetProductResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *DetailsServiceGetProductResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewGetProductResp()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *DetailsServiceGetProductResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetProduct_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *DetailsServiceGetProductResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *DetailsServiceGetProductResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DetailsServiceGetProductResult(%+v)", *p)
}

func (p *DetailsServiceGetProductResult) DeepEqual(ano *DetailsServiceGetProductResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *DetailsServiceGetProductResult) Field0DeepEqual(src *GetProductResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
