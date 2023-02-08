// Test the function that fetches all articles
package main

import "testing"

func TestGetAllArticles(t *testing.T) {

	alist := getAllArticles()

	// Check that the length of articles returned is the
	// smae as the length of the global variable holding the list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check that each membr is indentical
	for i, v := range alist {
		if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}
