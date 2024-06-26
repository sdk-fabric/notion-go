
// database automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


import "time"
type Database struct {
    Object string `json:"object"`
    Id string `json:"id"`
    CreatedTime time.Time `json:"created_time"`
    LastEditedTime time.Time `json:"last_edited_time"`
    Icon *DatabaseIcon `json:"icon"`
    Cover *DatabaseCover `json:"cover"`
    Url string `json:"url"`
    Title []DatabaseTitle `json:"title"`
    Description []DatabaseTitle `json:"description"`
}
