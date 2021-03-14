const PROTOCOL = window.location.protocol;
const BASE_PATH = window.location.host;
const REQUEST_URL = new URL('/api/v1/cards?name=', PROTOCOL + BASE_PATH);
const CARD_IMG_CLASS = "card-img";

async function queryCards(searchInput) {
    const res = await fetch(REQUEST_URL + searchInput.value);
    const responseJson = await res.json();
    const cards = responseJson.data;
    cardsWithEnglishScan = cards.filter(card => card.imageurlen != false);
    populateCardResults(cardsWithEnglishScan);
}


function populateCardResults(cards) {
    const cardResultsDiv = document.querySelector('[data-name="card-results"]');
    cardResultsDiv.innerHTML = "";
    for (const card of cards) {
        const cardDiv = document.createElement("div");
        cardDiv.classList.add(CARD_IMG_CLASS);
        cardDiv.style.backgroundImage = `url(${card.imageurlen})`;
        
        cardImgLink = document.createElement("a");
        cardImgLink.href = card.imageurlen;
        cardImgLink.target = "_blank";

        cardImgLink.appendChild(cardDiv);

        cardResultsDiv.appendChild(cardImgLink);
    }
}


function main() {
    const searchInput = document.getElementById("searchInput");
    searchInput.onchange = () => queryCards(searchInput);
}

main();