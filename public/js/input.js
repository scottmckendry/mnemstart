let inputSequence = [];
let leaderMode = false;

const mappingsArray = JSON.parse(
    document.getElementById("mappings").textContent,
);
const userSettings = JSON.parse(
    document.getElementById("userSettings").textContent,
);

const keymaps = parseMappings(mappingsArray);
let leaderKey = userSettings.LeaderKey;

if (leaderKey.length !== 1) {
    setStatus(
        "Invalid leader key. Please set a single character leader key in your settings. Using default leader key.",
    );
    leaderKey = " ";
}

document.addEventListener("keydown", (event) => {
    const key = event.key;
    setStatus("");

    // ignore keypresses when an input is focused
    const activeElement = document.activeElement;
    if (activeElement && activeElement.tagName.toLowerCase() === "input") {
        return;
    }

    if (key === leaderKey) {
        setStatus("Listening for key map...");
        leaderMode = true;
        inputSequence = [];
        return;
    }

    if (leaderMode) {
        inputSequence.push(key);

        const longestSequence = Math.max(
            ...Object.keys(keymaps).map((s) => s.length),
        );

        if (inputSequence.length > longestSequence) {
            inputSequence.shift();
        }

        const inputString = inputSequence.join("");
        setStatus(inputString);

        for (const mapping in keymaps) {
            if (inputString.endsWith(mapping)) {
                setStatus(mapping + " → " + keymaps[mapping]);
                window.location.href = keymaps[mapping];
                inputSequence = [];
                leaderMode = false;
                break;
            }
        }

        // If no match found, exit leader mode
        if (
            !Object.keys(keymaps).some((mapping) =>
                mapping.startsWith(inputString),
            )
        ) {
            setStatus("No matching key map found for " + inputString);
            leaderMode = false;
        }
    }
});

// default mappings
// 'esc' clears the input sequence and unfocuses any input fields
document.addEventListener("keydown", (event) => {
    if (event.key === "Escape") {
        event.preventDefault();
        inputSequence = [];
        leaderMode = false;
        document.activeElement.blur();
    }
});

// 'i' focuses the search input - i.e. 'insert mode'
document.addEventListener("keydown", (event) => {
    const activeElement = document.activeElement;
    if (activeElement && activeElement.tagName.toLowerCase() === "input") {
        return;
    }

    if (event.key === "i" && !leaderMode) {
        event.preventDefault();
        document.getElementById("search").focus();
        document.getElementById("search").value = "";
        document
            .getElementById("search")
            .addEventListener("keydown", (event) => {
                if (event.key === "Escape") {
                    document.getElementById("search").blur();
                }
                event.stopPropagation();
            });
    }
});

// Parses an array of mappings objects into a dictionary
function parseMappings(mappingsArray) {
    const mappings = {};
    mappingsArray.forEach((mapping) => {
        mappings[mapping["Keymap"]] = mapping["MapsTo"];
    });
    return mappings;
}

function setStatus(status) {
    document.getElementById("status").textContent = status;
}
