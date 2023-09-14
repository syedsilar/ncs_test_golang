// ui/main.js

document.addEventListener("DOMContentLoaded", function () {
    const userRegistrationForm = document.getElementById("user-registration-form");
    const userLoginForm = document.getElementById("user-login-form");
    const parkingAvailabilityDiv = document.getElementById("parking-availability");
    const refreshButton = document.getElementById("refresh-parking");
    const parkingSpotsDiv = document.getElementById("parking-spots");

    userRegistrationForm.addEventListener("submit", async function (event) {
        event.preventDefault();
        const username = document.getElementById("username").value;
        const password = document.getElementById("password").value;
    });

    userLoginForm.addEventListener("submit", async function (event) {
        event.preventDefault();
        const username = document.getElementById("login-username").value;
        const password = document.getElementById("login-password").value;
    });

    refreshButton.addEventListener("click", async function () {
    });
});
