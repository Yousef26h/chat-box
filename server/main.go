package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Hub struct {
}

func main() {

	// scanner := bufio.NewScanner(os.Stdin)

	// fmt.Println("Enter text (Press Ctrl+C to stop):")

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Printf("Captured: %s\n", line)
	// }

	// Optional: Serve a minimal HTML page to test the client locally
	http.HandleFunc("/", serveHome)

	fmt.Println("Chat server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("static/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

// Embedded HTML string for visual browser testing
const htmlTestPage = `
<!DOCTYPE html>
<html>
<head><title>Go Chat Test</title></head>
<body>
    <div id="output" style="height: 300px; overflow-y: scroll; border: 1px solid #ccc; padding: 10px;"></div>
    <input type="text" id="input" placeholder="Type a message..." style="width: 80%;" />
    <button onclick="send()">Send</button>

    <script>
        const ws = new WebSocket("ws://" + window.location.host + "/ws");
        const output = document.getElementById("output");
        const input = document.getElementById("input");

        ws.onmessage = function(event) {
            output.innerHTML += "<div>" + event.data + "</div>";
            output.scrollTop = output.scrollHeight;
        };

        function send() {
            if (input.value.trim() !== "") {
                ws.send(input.value);
                input.value = "";
            }
        }
        input.addEventListener("keypress", function(e) { if(e.key === "Enter") send(); });
    </script>
</body>
</html>
`
