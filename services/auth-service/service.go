package main

func AuthenticateUser(username, password string) bool {
    // Dummy authentication logic
    if username == "admin" && password == "password" {
        return true
    }
    return false
}
