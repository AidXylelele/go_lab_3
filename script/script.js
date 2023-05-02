import { diagonal, greenBtn, moveFigure } from "./processing-functions.js";
import { sendHTTPRequest } from "./request-api.js";
import { UrlService } from "./services.js";

const form = document.querySelector('form');
const textarea = document.querySelector('textarea');
const greenFrameBtn = document.querySelector(".gf-script");
const diagonalMoveBtn = document.querySelector(".dm-script");

greenFrameBtn.addEventListener("click", () => {
  const url = UrlService.create(greenBtn);

  sendHTTPRequest(url)
    .then((response) => console.log(response))
    .catch((error) => console.error(error));
});

form.addEventListener("submit", (e) => {
  e.preventDefault();
  const commandString = textarea.value.trim();
  const parsedCommands = UrlService.parseCommandString(commandString);
  const url = UrlService.create(parsedCommands);

  sendHTTPRequest(url)
    .then((response) => console.log(response))
    .catch((error) => console.error(error));
});

diagonalMoveBtn.addEventListener("click", async () => {
  const url = UrlService.create(diagonal);
  await sendHTTPRequest(url);
  await moveFigure();
});
