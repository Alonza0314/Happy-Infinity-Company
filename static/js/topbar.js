function topbarLanguage() {
    var link = document.getElementById("link");
    var linken = document.getElementById("link-en");
    
    if (link.style.display === "none") {
        link.style.display = "unset";
        linken.style.display = "none";
    } else {
        link.style.display = "none";
        linken.style.display = "unset";
    }
}
