

function tobytes(str) {
    var bytes = [];
    for(var i = 0; i < str.length; i++) {
        var char = str.charCodeAt(i);
        bytes.push(char >>> 8);
        bytes.push(char & 0xFF);
    }
    return bytes;
}

function bin2string(array) {
    var result = "";
    for (var i = 0; i < array.length; i++) {
      result += String.fromCharCode(parseInt(array[i], 2));
    }
    return result;
  }

console.log("script start!")

function handler(request, body, response) {
    console.log("handle")
    console.log(request.Path)
    console.log(body)

    console.log(JSON.parse(body).aa)

    data = tobytes(JSON.stringify({'aa': 'xxx'}))
    response.Header().Add('content-type', 'aaa')
    response.Write(data)
    response.WriteHeader(400)
}

