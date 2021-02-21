const PROTOCOL = window.location.protocol;
const BASE_PATH = window.location.host;
const REQUEST_URL = new URL('/api/v1/cards?name=', PROTOCOL + BASE_PATH);
const CARD_IMG_CLASS = "card-img";

async function queryCards(searchBox) {
    const res = await fetch(REQUEST_URL + searchBox.value);
    const cards = await res.json();
    populateCardResults(cards);
}


function populateCardResults(cards) {
    const cardResultsDiv = document.querySelector('[data-name="card-results"]');
    cardResultsDiv.innerHTML = "";
    for (const card of cards) {
        const cardDiv = document.createElement("div");
        cardDiv.classList.add(CARD_IMG_CLASS);
        cardDiv.style.backgroundImage = `url(${card.imageurlen})`;
        cardResultsDiv.appendChild(cardDiv);
    }
}


function main() {
    const searchBox = document.getElementById("searchBox");
    searchBox.onchange = () => queryCards(searchBox);
}

main();