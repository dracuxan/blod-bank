const url = "https://cat-fact.herokuapp.com/facts";

const factPara = document.querySelector("#fact");
const factbtn = document.querySelector("#btn");
factbtn.innerText = "Get Fact";
let i = 0;

const getFacts = async () => {
  let response = await fetch(url);
  console.log("Getting data...");
  let data = await response.json();
  factPara.innerText = data[i].text;
  factbtn.innerText = "New Fact";
  checkI();
};

const checkI = () => {
  if(i === 4) {
    i = 0;
  } else {
    i += 1;
  }
};

btn.addEventListener("click", getFacts);

