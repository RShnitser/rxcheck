let authToken = null;

document.body.addEventListener("login", function (e) {
  authToken = e.detail.token;
});

document.body.addEventListener("htmx:configRequest", function (e) {
  if (authToken === null) {
    return;
  }
  e.detail.headers["Authorization"] = `Bearer ${authToken}`;
});
