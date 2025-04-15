-- Migration: 002_seed_admin_user

-- Up migration
-- Insert a default admin user (password: admin123)
INSERT INTO users (external_id, username, email, password_hash, role)
VALUES 
  ('user-123', 'admin', 'admin@example.com', '$2a$10$zL.MmDQXIaQNgVLTj6Shs.Xs.R2f1QZn2qWbGa.EOOE3NwR9F5G8.', 'admin')
ON CONFLICT (username) DO NOTHING;

-- Down migration
DELETE FROM users WHERE username = 'admin'; 