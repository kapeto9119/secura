const AuditTrail = artifacts.require("AuditTrail");

module.exports = function (deployer) {
  // Deploy the AuditTrail contract
  deployer.deploy(AuditTrail);
}; 