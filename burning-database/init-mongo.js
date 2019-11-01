db.createUser({
    user: "andrii",
    pwd: "admin",
    roles: [
        {
            role: "readWrite",
            db: "burning-db"
        }
    ]
});