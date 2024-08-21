let inputSequence = [];
let leaderMode = false;

const userSettings = JSON.parse(
    document.getElementById("userSettings").textContent,
);

const keymaps = userSettings.Mappings;
const leaderKey = userSettings.LeaderKey;

document.addEventListener("keydown", (event) => {
    const key = event.key;

    // ignore keypresses when an input is focused
    const activeElement = document.activeElement;
    if (activeElement && activeElement.tagName.toLowerCase() === "input") {
        return;
    }

    if (key === leaderKey) {
        console.log("activating leader mode");
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
        console.log("inputString:", inputString);

        for (const mapping in keymaps) {
            if (inputString.endsWith(mapping)) {
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
            console.log("no match found, exiting leader mode");
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
