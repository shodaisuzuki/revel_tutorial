// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}


type tArticleApi struct {}
var ArticleApi tArticleApi


func (_ tArticleApi) GetArticles(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ArticleApi.GetArticles", args).URL
}

func (_ tArticleApi) GetArticle(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ArticleApi.GetArticle", args).URL
}

func (_ tArticleApi) PostArticle(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ArticleApi.PostArticle", args).URL
}

func (_ tArticleApi) PutArticle(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ArticleApi.PutArticle", args).URL
}

func (_ tArticleApi) DeleteArticle(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ArticleApi.DeleteArticle", args).URL
}


type tTestApi struct {}
var TestApi tTestApi


func (_ tTestApi) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestApi.Index", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


