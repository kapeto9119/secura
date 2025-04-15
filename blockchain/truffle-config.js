/**
 * Truffle configuration for the Secura blockchain integration
 */
module.exports = {
  // Configure networks
  networks: {
    // Development network (local)
    development: {
      host: "127.0.0.1",     // Localhost (default: none)
      port: 8545,            // Standard Ethereum port (default: none)
      network_id: "*",       // Any network (default: none)
    },
    // Ganache for local testing via docker-compose
    ganache: {
      host: "ganache",       // Docker service name
      port: 8545,
      network_id: "5777",
    },
  },

  // Set default mocha options
  mocha: {
    timeout: 100000
  },

  // Configure Solidity compiler
  compilers: {
    solc: {
      version: "0.8.17",      // Fetch exact version from solc-bin
      settings: {
        optimizer: {
          enabled: true,
          runs: 200
        },
      }
    }
  },

  // Configure plugins
  plugins: [
    'truffle-plugin-verify'
  ],

  // Configure contract paths
  contracts_directory: './contracts/',
  contracts_build_directory: './build/contracts/',
  migrations_directory: './migrations/',
}; 