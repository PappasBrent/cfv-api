// TODO: Revamp HTML to uses classes for styling, and data attributes for JS


const PROTOCOL = window.location.protocol;
const BASE_PATH = window.location.host;
const REQUEST_URL = new URL('/api/v1/cards?name=', PROTOCOL + BASE_PATH);

async function queryCards(searchBox) {
    const res = await fetch(REQUEST_URL + searchBox.value);
    const cards = await res.json();
    // console.table(resJson);
    populateCardResults(cards);
}


function populateCardResults(cards) {
    const cardResultDiv = document.querySelector(".card-results");
    // TODO: Add each card in the card results list to the cardResultsDiv
    // Will be a div with class card-img
}


function main() {
    const searchBox = document.getElementById("searchBox");
    searchBox.onchange = () => queryCards(searchBox);
}

main();