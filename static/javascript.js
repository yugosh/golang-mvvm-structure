// Fungsi untuk menyimpan data login ke localStorage
function saveLoginData(user) {
    localStorage.setItem('user', JSON.stringify(user));
}

// Fungsi untuk mengambil data login dari localStorage
function getLoginData() {
    return JSON.parse(localStorage.getItem('user'));
}

// Contoh penggunaan: Simpan data setelah login berhasil
function handleLoginResponse(response) {
    if (response.user) {
        saveLoginData(response.user);
        console.log('User data saved:', response.user);
    } else {
        console.error('Login failed:', response.message);
    }
}