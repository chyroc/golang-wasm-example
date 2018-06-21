function fetchAndInstantiate(url, importObject) {
    return fetch(url).then(response =>
        response.arrayBuffer()
    ).then(bytes =>
        WebAssembly.instantiate(bytes, importObject)
    ).then(results =>
        results.instance
    );
}

var go = new Go();
var mod = fetchAndInstantiate("./example.wasm", go.importObject);
window.onload = function () {
    mod.then(function (instance) {
        go.run(instance);
    });
};
