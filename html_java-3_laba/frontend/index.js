const backendLocation = "http://localhost:8080"
const apiLocation = "/api"
const commentsLocation = backendLocation + apiLocation + "/comments"
const usersLocation = backendLocation + apiLocation + "/users"

const userNameKey = "user_name"
const messageKey = "message"
const existKey = "exist"

let lastAuthorizedUser = null

window.onload = async () => {
    document.getElementById("addCommentButton").addEventListener("click", addCommentAction)

    let json = await getCommentsRequest()
    if (json === undefined) {
        return
    }

    json.forEach((comment) => {
        addCommentHtml(comment[userNameKey], comment[messageKey])
    })
}

async function getCommentsRequest() {
    let response = await fetch(commentsLocation)

    if (!response.ok) {
        console.warn("cannot get comments from db")
        return undefined
    }
    return await response.json()
}

async function addCommentRequest(name, message) {
    let response = await fetch(usersLocation+"/exist?" + new URLSearchParams({
        "user": name
    }))

    if (!response.ok) {
        console.warn("cannot check if user exist")
        return false
    }

    let json = await response.json()
    if (!json[existKey]) {
        response = await fetch(usersLocation,{
            method: "POST",
            body: JSON.stringify({
                "name": name
            })
        })

        if (!response.ok) {
            console.warn("cannot add user")
            return false
        }
    }

    response = await fetch(commentsLocation+"/exist?" + new URLSearchParams({
        "user_name": name,
        "message": message
    }))
    if (!response.ok) {
        console.warn("cannot check if comment exist")
        return false
    }

    json = await response.json()
    if (json[existKey]) {
        return false
    }

    response = await fetch(commentsLocation, {
        method: "POST",
        body: JSON.stringify({
            "user_name": name,
            "message": message
        })
    })

    if (!response.ok) {
        console.warn("cannot add comment")
        return false
    }
    return true
}

const addCommentHtml = (name, message) => {
    let newComment = document.createElement("P")
    newComment.innerHTML =
        "<div class='comment'>" +
        "<u class='comment_user'>" + name + "</u>" +
        "<br/>" + message +
        "<button class='delete_comment_button' style='visibility: hidden'>Удалить</button>"
    "</div>"
    newComment.firstChild.childNodes.item(3).addEventListener("click", deleteCommentAction)

    let commentsBlock = document.getElementById("commentsBlock")
    commentsBlock.appendChild(newComment)
}

async function addCommentAction() {
    let name = document.getElementById("nameInput").value
    let message = document.getElementById("messageInput").value

    if (!await addCommentRequest(name, message)) {
        return
    }

    let addCommentLabel = document.getElementById("addCommentLabel")
    if (lastAuthorizedUser === null) {
        addCommentLabel.innerHTML += ", " + name
    } else {
        let lastWord = addCommentLabel.innerHTML.split(" ").pop()
        addCommentLabel.innerHTML = addCommentLabel.innerHTML.replace(lastWord, name)
    }
    lastAuthorizedUser = name

    addCommentHtml(name, message)
    changeDeleteButtonVisibility()
}

async function deleteCommentRequest(name, message) {
    let response = await fetch(commentsLocation, {
        method: "DELETE",
        body: JSON.stringify({
            "user_name": name,
            "message": message
        })
    })

    if (!response.ok) {
        console.warn("cannot delete comment")
        return false
    }
    return true
}

const deleteCommentHtml = (element) => {
    element.remove()
}

async function deleteCommentAction() {
    let comment = event.target.parentElement
    let nodes = comment.childNodes

    if (! await deleteCommentRequest(nodes.item(0).innerHTML, nodes.item(2).data)) {
        return
    }

    deleteCommentHtml(comment)
}

const changeDeleteButtonVisibility = () => {
    let comments = document.getElementsByClassName("comment")

    for (let comment of comments) {
        let nodes = comment.childNodes

        if (nodes.item(0).innerHTML === lastAuthorizedUser) {
            nodes.item(3).style.visibility = "visible"
        } else {
            nodes.item(3).style.visibility = "hidden"
        }
    }
}
