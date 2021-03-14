const PROTOCOL = window.location.protocol;
const BASE_PATH = window.location.host;
const REQUEST_URL = new URL('/api/v1/cards?name=', PROTOCOL + BASE_PATH);
const CARD_IMG_CLASS = "card-img";

async function queryCards(searchInput) {
    const res = await fetch(REQUEST_URL + searchInput.value);
    const cards = await res.json();
    cardsWithEnglishScan = cards.filter(card => card.imageurlen != false)
    populateCardResults(cardsWithEnglishScan);
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
    const searchInput = document.getElementById("searchInput");
    searchInput.onchange = () => queryCards(searchInput);
}

main();