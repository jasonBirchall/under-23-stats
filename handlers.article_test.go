package main

// Test that a GET request to the home page returns a 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := setupRouter()

	r.GET("/", showIndexPage)