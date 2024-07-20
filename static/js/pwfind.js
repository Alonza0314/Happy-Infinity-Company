document.addEventListener("DOMContentLoaded", function() {
    const captcha = document.getElementById("captcha");
    const refresh = document.getElementById("refresh-captcha");
    let correctAuthCode = "";

    function refreshCaptcha() {
        const authUrl = captcha.dataset.authUrl;
        fetch(authUrl)
            .then(response => response.json())
            .then(data => {
                captcha.src = `${data.image}`;
                correctAuthCode = data.code;
            })
            .catch(error => console.error('Error fetching AuthCode:', error));
    }

    refreshCaptcha();

    refresh.addEventListener("click", function() {
        refreshCaptcha();
    });

    function checkPwfindNull(username, email, authcode) {
        return username !== "" && email !== "" && authcode !== "";
    }
    
    function checkPwfindUsername(username) {
        var regex = /^(?=.*\d)[a-z\d]+$/;
        return regex.test(username);
    }
    
    function checkPwfindEmail(email) {
        var regex = /\S+@\S+\.\S+/;
        return regex.test(email);
    }
    
    function checkPwfind() {
        var username = document.getElementById("username").value;
        var email = document.getElementById("email").value;
        var authcode = document.getElementById("captcha-code").value
    
        if (!checkPwfindNull(username, email, authcode)) {
            alert("blank field\n\n欄位空白");
            return false;
        }
    
        if (!checkPwfindUsername(username)) {
            alert("username can only and must contain lowercase English and numbers\nplease re-enter\n\n使用者名稱僅能且必須包含小寫英文及數字，請重新輸入");
            return false;
        }
    
        if (!checkPwfindEmail(email)) {
            alert("email format error\nplease re-enter\n\n電子郵件格式錯誤，請重新輸入");
            return false;
        }

        return true;
    }

    document.querySelector('.pwfind-box').addEventListener('submit', function(event) {
        event.preventDefault();

        if (!checkPwfind()) {
            return;
        }

        const captchaCodeInput = document.getElementById("captcha-code").value;

        if (correctAuthCode === captchaCodeInput) {
            this.submit();
        } else {
            alert("captcha code error\nplease re-enter\n\n驗證碼錯誤，請重新輸入");
            captchaCodeInput.value = "";
            refreshCaptcha();
            return
        }
    });
});

function getUrlParameter(name) {
    name = name.replace(/[\[]/, '\\[').replace(/[\]]/, '\\]');
    var regex = new RegExp('[\\?&]' + name + '=([^&#]*)');
    var results = regex.exec(location.search);
    return results === null ? '' : decodeURIComponent(results[1].replace(/\+/g, ' '));
}

var pwfindCondition = getUrlParameter('pwfind');
if (pwfindCondition !== '') {
    alert('fail: ' + pwfindCondition);
}
