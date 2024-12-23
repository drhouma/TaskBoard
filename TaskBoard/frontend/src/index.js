const form = document.getElementById("todo-form");
const input = document.getElementById("todo-input");
const todoLane = document.getElementById("todo-lane");
const inProgressLane = document.getElementById("in-progress-lane")
const doneLane =  document.getElementById("done-lane")
const user = localStorage.getItem("user")

const todoCategory = "todo"
const inProgressCategory = "in progress"
const doneCategory = "done"

window.onload = async () => {
    // prevent not logged user to use (idk default value)
    if (!user || user === "undefined" || user === null || user === "null" || user === "") {
        window.location.replace("./")
        return
    }

    let resp = await fetch("http://localhost:8080/api/tasks")
    if (!resp.ok) {
        window.alert("error while fetch all tasks")
    }

    let json = await resp.json()
    json.forEach((task) => {
        addTaskHtml(task["user"], task["description"], task["category"])
    })
}

form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const value = input.value;

    if (!value) return;

    if (await addTaskApi(user, value, todoCategory)) {
        addTaskHtml(user, value, todoCategory)
    } else {
        window.alert("error while fetch all tasks")
    }
    input.value = "";
});

form.addEventListener("reset", (e) => {
    e.preventDefault();

    localStorage.setItem("user", undefined)
    window.location.replace("./")
})

const addTaskApi = async (user, description, category) => {
    if (!user || !description) return

    let resp = await fetch("http://localhost:8080/api/tasks", {
        method: "POST",
        body: JSON.stringify({
            "user": user,
            "description": description,
            "category": category
        })
    })
    return resp.ok
}

const deleteTaskApi = async (user, description) => {
    if (!user || !description) return

    let resp = await fetch("http://localhost:8080/api/tasks?" + new URLSearchParams({
        "user": user,
        "description": description
    }), {
        method: "DELETE"
    })
    return resp.ok
}

const updateTaskApi = async (user, description, category) => {
    if (!user || !description || !category) return

    let resp = await fetch("http://localhost:8080/api/tasks?" + new URLSearchParams({
        "user": user,
        "description": description,
        "category": category,
    }), {
        method: "PATCH"
    })
    return resp.ok
}

const addTaskHtml = (author, description, group) => {
    if (!description || !author || !group) return

    const authorContainer = document.createElement("h4")
    authorContainer.innerText = author

    const bottomContainer = document.createElement("div")
    bottomContainer.classList.add("task-bottom")
    bottomContainer.appendChild(authorContainer)

    if (author === user) {
        const deleteButton = document.createElement("button")
        deleteButton.classList.add("delete-button")
        deleteButton.setAttribute("type", "submit")
        deleteButton.innerText = "delete"
        deleteButton.addEventListener("click", async () => {
            if (await deleteTaskApi(user, description)) {
                newTask.remove()
            } else {
                window.alert("cannot delete task")
            }
        })

        bottomContainer.appendChild(deleteButton)
    }

    const descriptionContainer = document.createElement("p")
    descriptionContainer.innerText = description;

    const newTask = document.createElement("div")
    newTask.classList.add("task")
    newTask.appendChild(descriptionContainer)
    newTask.appendChild(bottomContainer)

    if (user === author) {
        newTask.setAttribute("draggable", "true")
        newTask.addEventListener("dragstart", () => {
            newTask.classList.add("is-dragging");
        });
        newTask.addEventListener("dragend", () => {
            if (!updateTaskApi(user, description, newTask.parentElement.getAttribute("category"))) {
                window.alert("cannot change task category")
            }
            newTask.classList.remove("is-dragging");
        })
        newTask.style.cursor = "move"
    }

    switch (group) {
        case todoCategory: {
            todoLane.appendChild(newTask)
            break
        }
        case inProgressCategory: {
            inProgressLane.appendChild(newTask)
            break
        }
        case doneCategory: {
            doneLane.appendChild(newTask)
            break
        }
        default: {
            window.alert("unknown group")
            break
        }
    }
}