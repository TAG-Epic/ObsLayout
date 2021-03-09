let urlParams = new URLSearchParams(window.location.search);
let statType = urlParams.get("stat");
setInterval(() => {
    fetch("/stat?stat=" + statType).then(async (res) => {
        let container = document.getElementById("container");
        container.innerHTML = await res.text();
    })
}, 100);