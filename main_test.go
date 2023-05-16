package main

func Test(t *testing.T) {
	testURL = "http://" + getEnvValue("HOST") + ":" + getEnvValue("PORT")
}
