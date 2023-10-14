db.createCollection('users'); // essa linha não precisa já que estou fazendo o insert na linha 3

db.users.insertMany([
    {
        ID: '1',
        email: 'renato@likwi.com.br',
        password: '01_@exemplo@',
        name: 'renato',
        age: 18
    }
]);
