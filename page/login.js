
    window.onload = function () {
    const xhttp = new XMLHttpRequest();
    xhttp.withCredentials =true

    xhttp.open("GET", "http://localhost/api/login")
    xhttp.send();

    }
    function sendJSON() {
      let result = document.querySelector('.result');
      let username = document.querySelector('#username');
      let password = document.querySelector('#password');

      // Creating a XHR object
      let xhr = new XMLHttpRequest();
      let url = "http://localhost/api/login";

      // open a connection
      xhr.open("POST", url, false);

      // Set the request header i.e. which type of content you are sending
      xhr.setRequestHeader("Content-Type", "application/json");

      // Create a state change callback
      xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
          // Print received data from server
          result.innerHTML = this.responseText;
         window.location.href = "http://localhost/about.html"
        } else {

          result.innerHTML = this.responseText;
        }
      };

      // Converting JSON data to string
      var data = JSON.stringify({ "username": username.value, "password": password.value });

      // Sending data with the request
      xhr.send(data);
    }
