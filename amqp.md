## USER

* VerifyToken
> queue name: auth.verifyToken
> request: "token"
> response: {"model":vo.Authuser}

* GetUserNameById
> queue name: auth.getUserNameById
> request: [id1,id2...]
> response: {"model":{"id1":"username1","id2":"username2"}}

* SelectAdmin
> queue name: auth.selectAdmin
> request: 无参数
> response: {"model":{"id":id,"username":"xxx"...}}

* GetUserById
> queue name: auth.getUserById
> request: [id1,id2...]
> response: {"model":{"id1":{"name":"username1"},"id2":{"name":"username2"}}}

## CATEGORY

* GetCategoryNameById
> queue name: category.getCategoryNameById
> request: [id1,id2...]
> response: {"model":{"id1":"categoryname1","id2":"categoryname2"}}

* GetTagsByName
> queue name: tags.getTagsByName
> request: "tagsName"
> response: {"model":{"id":tags_id,"name":tags_name}}

* GetTagsByIds
> queue name: tags.getTagsByIds
> request: [id1,id2,...]
> response: {"models":[{"id":tags_id,"name":tags_name}...]}

* AddTags
> queue name: tags.addTags
> request: [vo.Tags,vo.Tags]
> response: {"models":[id1,id2...]}

## POSTS

* GetTagsIDAndCount
> queue name: posts.tags.getTagsIDAndCount
> request: 无参数
> response: 返回以tags_id为键，posts数为值的数组

* GetCategoryIDAndCount
> queue name: posts.getCategoryIDAndCount
> request: 无参数
> response: 返回以category_id为见，posts数为值的数组

## LOGS

* SaveLogs
> queue name: logs.saveLogs
> request: JSON(vo.AuthUserLog)

* GetParamGroupByCode
> queue name: logs.getParamGroupByCode
> request: "Code"
> response: {"models":[]vo.AuthUserLog{Paramter:"...",Code:"..."}}