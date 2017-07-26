package controllers

import (
    "github.com/revel/revel"
    "myapp/app/models"
)

type ArticleApi struct {
    *revel.Controller
}


func (c ArticleApi) GetArticles() revel.Result {

    // articleモデルを利用
    articles := []models.Article{}
    // Idが降順になるように取得
    DB.Order("id desc").Find(&articles)
    //この時点でarticleにはidが振られているのでそのまま返してあげます
    response := JsonResponse{}
    response.Response = articles // 結果を格納してあげる

    return c.RenderJSON(response)
}

func (c ArticleApi) GetArticle() revel.Result {

    // ルーティングで設定したurlに含まれる :id とかの部分はc.Params.Route.Getで取得
    id := c.Params.Route.Get("id")

    article := []models.Article{}
    // DB.Firstの第二引数にそのidを渡してあげると第一引数のモデルからidが一致するデータを検索してくれます
    DB.First(&article, id)

    response := JsonResponse{}
    response.Response = article

    return c.RenderJSON(response)
}

func (c ArticleApi) PostArticle() revel.Result {

    // articleモデルに値を格納
    article := &models.Article{
        // x-www-form-urlencodeで飛んできたデータはc.Params.Form.Getで受け取れます
        Title: c.Params.Form.Get("title"),
        Text: c.Params.Form.Get("text"),
    }

    // DBで保存
    DB.Create(article)

    response := JsonResponse{}

    response.Response = article

    return c.RenderJSON(response)
}

func (c ArticleApi) PutArticle() revel.Result {

    id := c.Params.Route.Get("id")

    article := &models.Article{}
    DB.First(&article, id)

    // DB.First()で返ってきたデータの中身を入れ直す
    article.Title = c.Params.Form.Get("title")
    article.Text = c.Params.Form.Get("text")

    // 入れ直したものをSave
    DB.Save(&article)

    response := JsonResponse{}
    response.Response = article

    return c.RenderJSON(response)
}

func (c ArticleApi) DeleteArticle() revel.Result {

    id := c.Params.Route.Get("id")

    article := []models.Article{}
    // 第二引数入れたidに一致するデータを第一引数のモデルから削除
    DB.Delete(&article, id)

    response := JsonResponse{}
    response.Response = "deleted. id: " + id

    return c.RenderJSON(response)
}
