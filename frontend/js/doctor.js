document.addEventListener("DOMContentLoaded", async function () {
  const token = localStorage.getItem("token");
  if (!token) {
    alert("Unauthorized! Please login.");
    window.location.href = "login.html";
    return;
  }

  try {
    const response = await fetch("http://localhost:8080/api/doctors", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Unauthorized or invalid token");
    }

    const doctors = await response.json();
    document.getElementById("doctors-list").innerHTML = doctors
      .map(
        (doctor) => `
            <p>${doctor.name} - ${doctor.specialization}</p>
        `
      )
      .join("");
  } catch (error) {
    console.error("Error fetching doctors:", error.message);
  }
});
