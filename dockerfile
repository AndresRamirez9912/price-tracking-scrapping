FROM golang:1.20

# Install Chromion as a dependency 
RUN apt-get update && apt-get install -y wget gnupg2 ca-certificates --no-install-recommends \
    && wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
    && echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list \
    && apt-get update \
    && apt-get install -y google-chrome-stable --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

# Create folder to store my project 
RUN mkdir /app

# Define the desired project address inside the container
WORKDIR /home/app 

# Copy the project from my pc
COPY . .

# Expose the port
EXPOSE 3002

# Run commands inside the container
RUN go get
CMD ["go", "run", "main.go"]
