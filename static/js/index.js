function Language() {
    var rh = document.getElementById("right-half");
    var rhen = document.getElementById("right-half-en");
    
    if (rh.style.display === "none") {
        rh.style.display = "flex";
        rhen.style.display = "none";
    } else {
        rh.style.display = "none";
        rhen.style.display = "flex";
    }
}
