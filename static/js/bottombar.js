function bottombarLanguage() {
    var refe = document.getElementById("refe");
    var refeen = document.getElementById("refe-en");
    
    if (refe.style.display === "none") {
        refe.style.display = "flex";
        refeen.style.display = "none";
    } else {
        refe.style.display = "none";
        refeen.style.display = "flex";
    }
}
