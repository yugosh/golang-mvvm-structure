const express = require('express');
const http = require('http');
const socketIo = require('socket.io');
const cors = require('cors'); // Import cors

const app = express();
const server = http.createServer(app);
const io = socketIo(server);

// Use CORS middleware to allow all origins
app.use(cors());
app.use(express.json()); // Middleware untuk parsing JSON

// Setup static file serving
app.use(express.static('public'));

let connectedClients = {};

// Handle Socket.IO connections
io.on('connection', (socket) => {
    console.log('New client connected:', socket.id);

    connectedClients[socket.id] = socket;

    socket.on('disconnect', () => {
        console.log('Client disconnected:', socket.id);

        delete connectedClients[socket.id];
    });

    // Emit message to the client
    socket.emit('new_message', { message: 'Welcome to Socket.IO server', socketId: socket.id });

    // Handle custom events
    socket.on('send_message', (msg) => {
        console.log('Received message from client:', msg);
        io.emit('new_message', { message: msg, socketId: socket.id });
    });
});

// Route to receive trigger from Golang
app.post('/trigger', (req, res) => {
    const { message, socketId } = req.body;  // Pastikan socketId diterima dari body
    if (socketId) {
        console.log('Received trigger from Golang:', message, 'with socketId:', socketId);

        // Emit message to the specific client with the matching socketId
        const clientSocket = io.sockets.sockets.get(socketId);
        if (clientSocket) {
            clientSocket.emit('new_message', { message, socketId });  // Emit message to the specific client
            console.log(`Message sent to client with socketId: ${socketId}`);
        } else {
            console.log('Socket ID not found:', socketId);
        }
    } else {
        console.log('No socketId provided');
    }
    res.sendStatus(200);
});

const PORT = process.env.PORT || 3000;
server.listen(PORT, () => console.log(`Server running on port ${PORT}`));
