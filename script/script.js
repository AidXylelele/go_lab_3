import { diagonal, greenBtn, moveFigure } from "./processing-functions.js";
import { sendHTTPRequest } from "./request-api.js";
import { UrlService } from "./services.js";

const greenFrameBtn = document.querySelector(".gf-script");
const diagonalMoveBtn = document.querySelector(".dm-script");

greenFrameBtn.addEventListener("click", () => {
  const url = UrlService.create(greenBtn);

  sendHTTPRequest(url)
    .then((response) => console.log(response))
    .catch((error) => console.error(error));
});

diagonalMoveBtn.addEventListener("click", async () => {
  const url = UrlService.create(diagonal);
  await sendHTTPRequest(url);
  await moveFigure();
});
