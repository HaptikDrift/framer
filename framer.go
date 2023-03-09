/*
  The following project is a proof of concept:
  Project Title: framer
  Goal or Aim:
   * To frame an external application in a webpage served by a local web server.
  ToDo:
   -
	written by Haptik Drift
	<haptikdrift@gmail.com>
*/


package main

/* All imports needed in the main function */
import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

/*
	###
	# APPLICATION FILESYSTEM
	###
*/
//go:embed assets/*
var assets embed.FS

/*
	###
	# Main Function
	###
*/

func main() {
	assetsFs := http.FileServer(http.FS(assets)) // Embed my assets folder into the applications FS (File System)
	http.Handle("/assets/", assetsFs)            // Set a URL where the assets can be found in the File System
	http.HandleFunc("/", framer)                 // Root Application/Home Page
	http.HandleFunc("/f", getframer)             // For GET Method framing
	fmt.Println("Starting server for testing HTTP POST ...\nhttp://127.0.0.1:8085")
	if err := http.ListenAndServe("127.0.0.1:8085", nil); err != nil {
		log.Fatal(err)
	}
}

/*	########################################################################################################	*/
/*
	###
	# APPLICATION FUNCTIONS
	###
*/
// Home Page
func framer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		render_page := page_header + page_body_end_submit
		fmt.Fprintf(w, `%v`, render_page)
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, `ParseForm() err: %v`, err)
			return
		}

		site := r.FormValue("frame")
		app := site
		render_page := page_header + open_frame + app + close_frame
		fmt.Fprintf(w, `%v`, render_page)
	default:
		fmt.Fprintf(w, `Sorry, only GET and POST methods are supported.`)
	}
}

// Framed Site page ... If framing worked.
func getframer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, `ParseForm() err: %v`, err)
			return
		}

		site := r.FormValue("frame")
		app := site
		render_page := page_header + open_frame + app + close_frame
		fmt.Fprintf(w, `%v`, render_page)
	default:
		fmt.Fprintf(w, `Sorry, only GET and POST methods are supported.`)
	}
}

/*
	###
	# HTML Page Constructors
	###
*/

/* The main page configuration */
var page_header string = `<html>
<style>
body {
	background: #aaa;
	display: flex;
	flex-wrap: wrap;
}

.module {
	background: white;
	border: 6px solid #008010;
	font-family: Arial, sans-serif;
	border-radius: 50px;
	margin: 2%;
	margin-right: 3%;
	width: 100%;
	padding: 1rem;
	> h1 {
		padding: 3 1rem;

	}
	> p {
	padding: 3 1rem;
	text-align: center;
	}
}

frame {
	padding: 3 9rem;
	border: 50px solid #000;
	display:block;
	margin: 0 auto;
}
.stripe {
	background-image: url("assets/256c9dc132d0a38a18f5e78e6c1b9d7c.png");
	background-repeat: no-repeat;
	background-position: center;
	margin: 2%;
	height: 120px;
}

.stripe2 {
	background-image: url("assets/a557780ed4d850b2a8449634f3ba87ea.png");
	background-repeat: no-repeat;
	background-position: center;
	margin: 2%;
	height: 120px;
}
</style>`

/* This is the tail of the main page configuration */
var page_body_end_submit string = `
	<title>Cross-Site Framing Demo: </title>
	<center>
	<body>
	<div class="module">
			<h1 class="stripe"> </h1>
			<p>You can submite a URL in the box below or you can use the get function and send the URL as:<br><pre>http://127.0.0.1:8085/f?frame=&lt;hostname&gt;</pre><p>Or pass the URL to GoWitness to take a screen-shot of the framed application.</p><pre>gowitness single --disable-db --delay 2 --fullpage 'http://127.0.0.1:8085/f?frame=&lt;hostname&gt;'</pre></p>
			<form action="/" method="post">
				<label>Site to Frame</label>
				<input type="text" name="frame" class="form-control" style="width: 849px;">
			  <!-- End Label and Textarea Field -->
				<button type="submit" class="btn btn-primary">Submit</button>
			  </div>
			<!-- End Label and Text Field -->
			</form>
		</div>
	<br>   
	</body>
	</center>
</html>`

var open_frame string = `
	<title>Cross-Site Framing Demo: </title>
		<center>
		<div class="module">
				<h1 class="stripe2"> </h1>
				<p> If web content is displayed below, the site or application is <b>VULNERABLE</b> to clickjacking/cross-site framing attacks. </p>
				<h1><div id="alertdiv"></div></h1>
				<iframe id="iframe" src="`

var close_frame string = `" width="1300" height="600" align="middle"></iframe>
		<script>
		var clicks = 0;
		var iframewatcher = setInterval(function(){
			var activeE = document.activeElement;
			if(activeE && activeE.tagName == 'IFRAME'){
				clicks += 1;
				alertdiv.textContent = clicks;
			}
			clearInterval(iframewatcher);
		}, 100);
		</script>
		</div>
		<br>
</html>`
