
// block_pdf automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


type BlockPDF struct {
    Object string `json:"object"`
    Id string `json:"id"`
    Parent *ParentId `json:"parent"`
    Type string `json:"type"`
    CreatedTime string `json:"created_time"`
    CreatedBy *User `json:"created_by"`
    LastEditedTime string `json:"last_edited_time"`
    LastEditedBy *User `json:"last_edited_by"`
    Archived bool `json:"archived"`
    InTrash bool `json:"in_trash"`
    HasChildren bool `json:"has_children"`
    Pdf *PDF `json:"pdf"`
}

