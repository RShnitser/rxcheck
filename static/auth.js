let authToken = null;

document.body.addEventListener("login", function (e) {
  console.log(e.details.value);
  authToken = e.details.value;
});

document.body.addEventListener("htmx:configRequest", function (e) {
  if (authToken === null) {
    return;
  }

  e.details.headers["Authorization"] = `Bearer ${token}`;
});
