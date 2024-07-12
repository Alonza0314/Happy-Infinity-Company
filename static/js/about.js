function Language() {
    var c = document.getElementById("container");
    var cen = document.getElementById("container-en");

    var aboutdescribe = document.getElementById("about-describe");
    var historydescribe = document.getElementById("history-describe");
    var membersdescribe = document.getElementById("members-describe");

    var aboutdescribeen = document.getElementById("about-describe-en");
    var historydescribeen = document.getElementById("history-describe-en");
    var membersdescribeen = document.getElementById("members-describe-en");

    if (c.style.display === "none") {
        c.style.display = "flex";
        cen.style.display = "none";
        if (historydescribeen.style.display !== "block" && membersdescribeen.style.display !== "block") {
            aboutdescribeen.style.display = "none";
            aboutdescribe.style.display = "block";
            historydescribe.style.display = "none";
            membersdescribe.style.display = "none";
        } else if (aboutdescribeen.style.display !== "block" && membersdescribeen.style.display !== "block") {
            historydescribeen.style.display = "none";
            aboutdescribe.style.display = "none";
            historydescribe.style.display = "block";
            membersdescribe.style.display = "none";
        } else {
            membersdescribeen.style.display = "none";
            aboutdescribe.style.display = "none";
            historydescribe.style.display = "none";
            membersdescribe.style.display = "block";
        }
    } else {
        c.style.display = "none";
        cen.style.display = "flex";
        if (historydescribe.style.display !== "block" && membersdescribe.style.display !== "block") {
            aboutdescribe.style.display = "none";
            aboutdescribeen.style.display = "block";
            historydescribeen.style.display = "none";
            membersdescribeen.style.display = "none";
        } else if (aboutdescribe.style.display !== "block" && membersdescribe.style.display !== "block") {
            historydescribe.style.display = "none";
            aboutdescribeen.style.display = "none";
            historydescribeen.style.display = "block";
            membersdescribeen.style.display = "none";
        } else {
            membersdescribe.style.display = "none";
            aboutdescribeen.style.display = "none";
            historydescribeen.style.display = "none";
            membersdescribeen.style.display = "block";
        }
    }
}

document.addEventListener('DOMContentLoaded', function() {
    var about = document.getElementById("about");
    var history = document.getElementById("history");
    var members = document.getElementById("members");

    var abouten = document.getElementById("about-en");
    var historyen = document.getElementById("history-en");
    var membersen = document.getElementById("members-en");

    var aboutdescribe = document.getElementById("about-describe");
    var historydescribe = document.getElementById("history-describe");
    var membersdescribe = document.getElementById("members-describe");

    var aboutdescribeen = document.getElementById("about-describe-en");
    var historydescribeen = document.getElementById("history-describe-en");
    var membersdescribeen = document.getElementById("members-describe-en");
    
    about.addEventListener('click', function() {
        aboutdescribe.style.display = "block";
        historydescribe.style.display = "none";
        membersdescribe.style.display = "none";
    });
    history.addEventListener('click', function() {
        aboutdescribe.style.display = "none";
        historydescribe.style.display = "block";
        membersdescribe.style.display = "none";
    });
    members.addEventListener('click', function() {
        aboutdescribe.style.display = "none";
        historydescribe.style.display = "none";
        membersdescribe.style.display = "block";
    });

    abouten.addEventListener('click', function() {
        aboutdescribeen.style.display = "block";
        historydescribeen.style.display = "none";
        membersdescribeen.style.display = "none";
    });
    historyen.addEventListener('click', function() {
        aboutdescribeen.style.display = "none";
        historydescribeen.style.display = "block";
        membersdescribeen.style.display = "none";
    });
    membersen.addEventListener('click', function() {
        aboutdescribeen.style.display = "none";
        historydescribeen.style.display = "none";
        membersdescribeen.style.display = "block";
    });
});