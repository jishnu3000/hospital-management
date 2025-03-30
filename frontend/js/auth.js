document.addEventListener("DOMContentLoaded", function () {
  const loginForm = document.getElementById("login-form");
  if (loginForm) {
    loginForm.addEventListener("submit", async function (event) {
      event.preventDefault();

      const email = document.getElementById("email").value;
      const password = document.getElementById("password").value;
      const errorMessage = document.getElementById("error-message");

      errorMessage.textContent = "";
      errorMessage.classList.remove("visible");

      try {
        const response = await fetch("http://localhost:8080/api/login", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ email, password }),
        });

        const data = await response.json();

        if (!response.ok) {
          throw new Error(data.error || "Invalid credentials");
        }

        localStorage.setItem("token", data.token);
        window.location.href = "dashboard.html";
      } catch (error) {
        errorMessage.textContent = error.message;
        errorMessage.classList.add("visible"); 
      }
    });
  }
});

const registerForm = document.getElementById("register-form");
if (registerForm) {
  registerForm.addEventListener("submit", async function (event) {
    event.preventDefault();

    const name = document.getElementById("name").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const errorMessage = document.getElementById("register-error-message");

    errorMessage.textContent = "";
    errorMessage.classList.remove("visible");

    try {
      const response = await fetch("http://localhost:8080/api/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, email, password }),
      });

      if (!response.ok) {
        const data = await response.json().catch(() => null);
        throw new Error(data?.error || "Registration failed");
      }

      // Redirect to login page after successful registration
      alert("Registration successful! Please login.");
      window.location.href = "login.html";
    } catch (error) {
      errorMessage.textContent = error.message;
      errorMessage.classList.add("visible"); 
    }
  });
}

// Logout function
window.logout = function () {
  localStorage.removeItem("token"); // Remove token
  window.location.href = "index.html";
};
