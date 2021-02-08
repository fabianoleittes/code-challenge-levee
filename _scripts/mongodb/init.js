db = db.getSiblingDB('levee');

db.createUser({
    user: 'levee',
    pwd: 'levee',
    roles: [
        {
            role: 'root',
            db: 'admin',
        },
    ],
});

db.createCollection('jobs');
