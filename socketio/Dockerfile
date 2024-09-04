# Menggunakan Node.js versi terbaru
FROM node:14

# Set working directory
WORKDIR /usr/src/app

# Copy package.json dan package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy seluruh source code ke working directory
COPY . .

# Expose port untuk Socket.IO (3000 atau port lain yang digunakan)
EXPOSE 3000

# Perintah untuk menjalankan aplikasi
CMD ["npm", "start"]
