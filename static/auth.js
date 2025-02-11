let authToken = null;

document.body.addEventListener("logIn", function (e) {
  authToken = e.details.value;
});

document.body.addEventListener("htmx:configRequest", function (e) {
  if (authToken === null) {
    return;
  }

  e.details.headers["Authorization"] = `Bearer ${token}`;
});
