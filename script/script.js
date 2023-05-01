const greenFrameBtn = document.querySelector(".gf-script");
const diagonalMoveBtn = document.querySelector(".dm-script");
const screenSize = 800;
const figureRadius = 300;
const startPos = figureRadius;
const finishPos = screenSize - figureRadius;
const maxDist = finishPos - startPos;
const step = 20;
const interval = 10000;
let pos = startPos;
let dist = 0;

const config = {
  greenBtn: {
    green: null,
    bgrect: [200, 200, 600, 600],
    update: null,
  },
  diagonal: {
    figure: [500, 500],
    update: null,
  },
  moveRight: {
    move: [0, 0],
  },
  moveLeft: {
    move: [0, 0],
  },
};

class UrlService {
  static url = "http://127.0.0.1:17000/?cmd=";

  static create = (options) => {
    let newUrl = this.url;
    for (const action in options) {
      newUrl += `${action}`;
      const arrayOfArgs = options[action];
      if (arrayOfArgs) newUrl += " " + arrayOfArgs.join(" ");
      newUrl += ",";
    }
    return newUrl.slice(0, -1);
  };
}

const moveFigure = async () => {
 
  await sendHTTPRequest(UrlService.create(config.diagonal));

 
};

const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

const sendHTTPRequest = async (url) => {
  try {
    const response = await fetch(url);
    if (response.ok) return response.text();
    throw new Error(`Request failed with status ${response.status}`);
  } catch (error) {
    console.error(error);
  }
};

greenFrameBtn.addEventListener("click", () => {
  const url = UrlService.create(config.greenBtn);

  sendHTTPRequest(url)
    .then((response) => console.log(response))
    .catch((error) => console.error(error));
});

diagonalMoveBtn.addEventListener("click", async () => {
  const url = UrlService.create(config.diagonal);
  await sendHTTPRequest(url);
  await moveFigure()
});
