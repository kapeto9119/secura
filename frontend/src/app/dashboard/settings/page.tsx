'use client';

import { useState } from 'react';

// Tab component for settings navigation
const SettingsTab = ({ 
  name, 
  current, 
  setCurrent 
}: { 
  name: string; 
  current: string; 
  setCurrent: (name: string) => void;
}) => {
  return (
    <button
      onClick={() => setCurrent(name)}
      className={`px-3 py-2 text-sm font-medium rounded-md ${
        current === name
          ? 'bg-blue-100 text-blue-700'
          : 'text-gray-500 hover:text-gray-700'
      }`}
    >
      {name}
    </button>
  );
};

// Form section component
const FormSection = ({ title, description, children }: { 
  title: string; 
  description: string; 
  children: React.ReactNode;
}) => {
  return (
    <div className="mt-10 divide-y divide-gray-200">
      <div className="space-y-1">
        <h3 className="text-lg font-medium leading-6 text-gray-900">{title}</h3>
        <p className="max-w-2xl text-sm text-gray-500">{description}</p>
      </div>
      <div className="mt-6 pt-6">
        {children}
      </div>
    </div>
  );
};

// Toggle switch component
const Toggle = ({ 
  enabled, 
  setEnabled, 
  label, 
  description 
}: { 
  enabled: boolean; 
  setEnabled: (enabled: boolean) => void; 
  label: string; 
  description?: string;
}) => {
  return (
    <div className="flex items-center justify-between">
      <div>
        <span className="text-sm font-medium text-gray-900">{label}</span>
        {description && (
          <p className="text-sm text-gray-500">{description}</p>
        )}
      </div>
      <button
        type="button"
        onClick={() => setEnabled(!enabled)}
        className={`${
          enabled ? 'bg-blue-600' : 'bg-gray-200'
        } relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2`}
      >
        <span className="sr-only">Toggle {label}</span>
        <span
          className={`${
            enabled ? 'translate-x-5' : 'translate-x-0'
          } pointer-events-none relative inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out`}
        >
          <span
            className={`${
              enabled
                ? 'opacity-0 duration-100 ease-out'
                : 'opacity-100 duration-200 ease-in'
            } absolute inset-0 flex h-full w-full items-center justify-center transition-opacity`}
            aria-hidden="true"
          >
            <svg className="h-3 w-3 text-gray-400" fill="none" viewBox="0 0 12 12">
              <path
                d="M4 8l2-2m0 0l2-2M6 6L4 4m2 2l2 2"
                stroke="currentColor"
                strokeWidth={2}
                strokeLinecap="round"
                strokeLinejoin="round"
              />
            </svg>
          </span>
          <span
            className={`${
              enabled
                ? 'opacity-100 duration-200 ease-in'
                : 'opacity-0 duration-100 ease-out'
            } absolute inset-0 flex h-full w-full items-center justify-center transition-opacity`}
            aria-hidden="true"
          >
            <svg className="h-3 w-3 text-blue-600" fill="currentColor" viewBox="0 0 12 12">
              <path d="M3.707 5.293a1 1 0 00-1.414 1.414l1.414-1.414zM5 8l-.707.707a1 1 0 001.414 0L5 8zm4.707-3.293a1 1 0 00-1.414-1.414l1.414 1.414zm-7.414 2l2 2 1.414-1.414-2-2-1.414 1.414zm3.414 2l4-4-1.414-1.414-4 4 1.414 1.414z" />
            </svg>
          </span>
        </span>
      </button>
    </div>
  );
};

// Input field component
const InputField = ({
  id,
  label,
  type = 'text',
  placeholder,
  value,
  onChange,
  helpText,
  required = false
}: {
  id: string;
  label: string;
  type?: string;
  placeholder?: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  helpText?: string;
  required?: boolean;
}) => {
  return (
    <div className="sm:grid sm:grid-cols-3 sm:gap-4 sm:items-start sm:pt-5">
      <label htmlFor={id} className="block text-sm font-medium text-gray-700 sm:mt-px sm:pt-2">
        {label} {required && <span className="text-red-500">*</span>}
      </label>
      <div className="mt-1 sm:mt-0 sm:col-span-2">
        <input
          type={type}
          name={id}
          id={id}
          value={value}
          onChange={onChange}
          placeholder={placeholder}
          className="max-w-lg block w-full shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:max-w-xs sm:text-sm border-gray-300 rounded-md"
          required={required}
        />
        {helpText && <p className="mt-2 text-sm text-gray-500">{helpText}</p>}
      </div>
    </div>
  );
};

// Select field component
const SelectField = ({
  id,
  label,
  value,
  onChange,
  options,
  helpText,
  required = false
}: {
  id: string;
  label: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  options: { value: string; label: string }[];
  helpText?: string;
  required?: boolean;
}) => {
  return (
    <div className="sm:grid sm:grid-cols-3 sm:gap-4 sm:items-start sm:pt-5">
      <label htmlFor={id} className="block text-sm font-medium text-gray-700 sm:mt-px sm:pt-2">
        {label} {required && <span className="text-red-500">*</span>}
      </label>
      <div className="mt-1 sm:mt-0 sm:col-span-2">
        <select
          id={id}
          name={id}
          value={value}
          onChange={onChange}
          className="max-w-lg block focus:ring-blue-500 focus:border-blue-500 w-full shadow-sm sm:max-w-xs sm:text-sm border-gray-300 rounded-md"
          required={required}
        >
          {options.map((option) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </select>
        {helpText && <p className="mt-2 text-sm text-gray-500">{helpText}</p>}
      </div>
    </div>
  );
};

// Account settings tab
const AccountSettings = () => {
  const [formData, setFormData] = useState({
    email: 'john.doe@example.com',
    name: 'John Doe',
    timezone: 'America/New_York'
  });

  const [twoFactorEnabled, setTwoFactorEnabled] = useState(false);
  const [notificationsEnabled, setNotificationsEnabled] = useState(true);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value
    }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // Save account settings
    console.log('Saving account settings:', { ...formData, twoFactorEnabled, notificationsEnabled });
    alert('Account settings saved successfully!');
  };

  return (
    <form onSubmit={handleSubmit}>
      <FormSection
        title="Profile Information"
        description="Update your account information and email address."
      >
        <div className="space-y-6">
          <InputField
            id="name"
            label="Name"
            value={formData.name}
            onChange={handleChange}
            required
          />
          <InputField
            id="email"
            label="Email"
            type="email"
            value={formData.email}
            onChange={handleChange}
            required
          />
          <SelectField
            id="timezone"
            label="Timezone"
            value={formData.timezone}
            onChange={handleChange}
            options={[
              { value: 'America/New_York', label: 'Eastern Time (US & Canada)' },
              { value: 'America/Chicago', label: 'Central Time (US & Canada)' },
              { value: 'America/Denver', label: 'Mountain Time (US & Canada)' },
              { value: 'America/Los_Angeles', label: 'Pacific Time (US & Canada)' },
              { value: 'Europe/London', label: 'London' },
              { value: 'Europe/Paris', label: 'Paris' },
              { value: 'Asia/Tokyo', label: 'Tokyo' }
            ]}
          />
        </div>
      </FormSection>

      <FormSection
        title="Security"
        description="Manage your account security settings."
      >
        <div className="space-y-6">
          <Toggle
            enabled={twoFactorEnabled}
            setEnabled={setTwoFactorEnabled}
            label="Two-factor authentication"
            description="Add an extra layer of security to your account."
          />
          <div className="pt-5">
            <button
              type="button"
              className="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
              Change Password
            </button>
          </div>
        </div>
      </FormSection>

      <FormSection
        title="Notifications"
        description="Configure how you receive notifications."
      >
        <Toggle
          enabled={notificationsEnabled}
          setEnabled={setNotificationsEnabled}
          label="Email notifications"
          description="Receive email notifications for important updates and activity."
        />
      </FormSection>

      <div className="pt-6 divide-y divide-gray-200">
        <div className="mt-4 pt-4 flex justify-end space-x-3">
          <button
            type="button"
            className="bg-white py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Cancel
          </button>
          <button
            type="submit"
            className="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Save Changes
          </button>
        </div>
      </div>
    </form>
  );
};

// API settings tab
const APISettings = () => {
  const [apiToken, setApiToken] = useState('••••••••••••••••••••••••••••••');
  const [rateLimitEnabled, setRateLimitEnabled] = useState(true);
  const [webhookEnabled, setWebhookEnabled] = useState(false);
  const [formData, setFormData] = useState({
    webhookUrl: '',
    rateLimit: '100'
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value
    }));
  };

  const generateNewToken = () => {
    // In a real app, this would call an API to generate a new token
    const newToken = Math.random().toString(36).substring(2, 15) + 
                    Math.random().toString(36).substring(2, 15);
    setApiToken(newToken);
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log('Saving API settings:', { ...formData, rateLimitEnabled, webhookEnabled });
    alert('API settings saved successfully!');
  };

  return (
    <form onSubmit={handleSubmit}>
      <FormSection
        title="API Keys"
        description="Manage your API keys that allow access to our services."
      >
        <div className="space-y-6">
          <div className="sm:grid sm:grid-cols-3 sm:gap-4 sm:items-start sm:pt-5">
            <label className="block text-sm font-medium text-gray-700 sm:mt-px sm:pt-2">
              API Token
            </label>
            <div className="mt-1 sm:mt-0 sm:col-span-2">
              <div className="flex max-w-lg rounded-md shadow-sm">
                <input
                  type="text"
                  value={apiToken}
                  disabled
                  className="flex-1 max-w-xs block w-full rounded-md border-gray-300 text-gray-500 bg-gray-50 sm:text-sm"
                />
                <button
                  type="button"
                  onClick={generateNewToken}
                  className="ml-3 inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                >
                  Generate New
                </button>
              </div>
              <p className="mt-2 text-sm text-gray-500">
                Keep this token secure. You cannot recover it once you leave this page.
              </p>
            </div>
          </div>
        </div>
      </FormSection>

      <FormSection
        title="Rate Limiting"
        description="Configure API rate limiting to prevent abuse."
      >
        <div className="space-y-6">
          <Toggle
            enabled={rateLimitEnabled}
            setEnabled={setRateLimitEnabled}
            label="Enable rate limiting"
            description="Limit the number of requests that can be made to the API."
          />
          
          {rateLimitEnabled && (
            <InputField
              id="rateLimit"
              label="Rate limit (requests per minute)"
              type="number"
              value={formData.rateLimit}
              onChange={handleChange}
              helpText="Maximum number of requests allowed per minute."
              required
            />
          )}
        </div>
      </FormSection>

      <FormSection
        title="Webhooks"
        description="Configure webhooks to receive real-time updates."
      >
        <div className="space-y-6">
          <Toggle
            enabled={webhookEnabled}
            setEnabled={setWebhookEnabled}
            label="Enable webhooks"
            description="Receive notifications via webhooks when events occur."
          />
          
          {webhookEnabled && (
            <InputField
              id="webhookUrl"
              label="Webhook URL"
              type="url"
              value={formData.webhookUrl}
              onChange={handleChange}
              helpText="The URL where webhook events will be sent."
              required
            />
          )}
        </div>
      </FormSection>

      <div className="pt-6 divide-y divide-gray-200">
        <div className="mt-4 pt-4 flex justify-end space-x-3">
          <button
            type="button"
            className="bg-white py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Cancel
          </button>
          <button
            type="submit"
            className="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Save Changes
          </button>
        </div>
      </div>
    </form>
  );
};

// Security & Privacy settings tab
const SecuritySettings = () => {
  const [auditLogsEnabled, setAuditLogsEnabled] = useState(true);
  const [ipWhitelistEnabled, setIpWhitelistEnabled] = useState(false);
  const [dataRetentionDays, setDataRetentionDays] = useState('90');
  const [whitelistedIps, setWhitelistedIps] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log('Saving Security settings:', { 
      auditLogsEnabled, 
      ipWhitelistEnabled, 
      dataRetentionDays,
      whitelistedIps: whitelistedIps.split('\n')
    });
    alert('Security settings saved successfully!');
  };

  return (
    <form onSubmit={handleSubmit}>
      <FormSection
        title="Audit Logs"
        description="Configure how audit logs are collected and stored."
      >
        <div className="space-y-6">
          <Toggle
            enabled={auditLogsEnabled}
            setEnabled={setAuditLogsEnabled}
            label="Enable audit logs"
            description="Record user activities for security and compliance purposes."
          />
          
          {auditLogsEnabled && (
            <div className="sm:grid sm:grid-cols-3 sm:gap-4 sm:items-start sm:pt-5">
              <label htmlFor="dataRetentionDays" className="block text-sm font-medium text-gray-700 sm:mt-px sm:pt-2">
                Data retention period (days)
              </label>
              <div className="mt-1 sm:mt-0 sm:col-span-2">
                <select
                  id="dataRetentionDays"
                  name="dataRetentionDays"
                  value={dataRetentionDays}
                  onChange={(e) => setDataRetentionDays(e.target.value)}
                  className="max-w-lg block focus:ring-blue-500 focus:border-blue-500 w-full shadow-sm sm:max-w-xs sm:text-sm border-gray-300 rounded-md"
                >
                  <option value="30">30 days</option>
                  <option value="60">60 days</option>
                  <option value="90">90 days</option>
                  <option value="180">180 days</option>
                  <option value="365">365 days</option>
                </select>
                <p className="mt-2 text-sm text-gray-500">
                  Audit logs will be automatically deleted after this period.
                </p>
              </div>
            </div>
          )}
        </div>
      </FormSection>

      <FormSection
        title="IP Access Controls"
        description="Restrict access to your account based on IP address."
      >
        <div className="space-y-6">
          <Toggle
            enabled={ipWhitelistEnabled}
            setEnabled={setIpWhitelistEnabled}
            label="Enable IP whitelist"
            description="Only allow access from specific IP addresses."
          />
          
          {ipWhitelistEnabled && (
            <div className="sm:grid sm:grid-cols-3 sm:gap-4 sm:items-start sm:pt-5">
              <label htmlFor="whitelistedIps" className="block text-sm font-medium text-gray-700 sm:mt-px sm:pt-2">
                Whitelisted IPs
              </label>
              <div className="mt-1 sm:mt-0 sm:col-span-2">
                <textarea
                  id="whitelistedIps"
                  name="whitelistedIps"
                  rows={4}
                  value={whitelistedIps}
                  onChange={(e) => setWhitelistedIps(e.target.value)}
                  className="max-w-lg shadow-sm block w-full focus:ring-blue-500 focus:border-blue-500 sm:text-sm border border-gray-300 rounded-md"
                  placeholder="Enter one IP address per line"
                />
                <p className="mt-2 text-sm text-gray-500">
                  Enter one IP address per line. Use CIDR notation for IP ranges (e.g., 192.168.1.0/24).
                </p>
              </div>
            </div>
          )}
        </div>
      </FormSection>

      <div className="pt-6 divide-y divide-gray-200">
        <div className="mt-4 pt-4 flex justify-end space-x-3">
          <button
            type="button"
            className="bg-white py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Cancel
          </button>
          <button
            type="submit"
            className="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Save Changes
          </button>
        </div>
      </div>
    </form>
  );
};

export default function SettingsPage() {
  const [currentTab, setCurrentTab] = useState('Account');
  const tabs = ['Account', 'API', 'Security & Privacy'];
  
  return (
    <div className="p-6">
      <div className="pb-5 border-b border-gray-200 sm:pb-0">
        <h1 className="text-2xl font-semibold leading-7 text-gray-900">Settings</h1>
        <div className="mt-3 sm:mt-4">
          <div className="sm:hidden">
            <label htmlFor="current-tab" className="sr-only">
              Select a tab
            </label>
            <select
              id="current-tab"
              name="current-tab"
              className="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm rounded-md"
              value={currentTab}
              onChange={(e) => setCurrentTab(e.target.value)}
            >
              {tabs.map((tab) => (
                <option key={tab}>{tab}</option>
              ))}
            </select>
          </div>
          <div className="hidden sm:block">
            <div className="border-b border-gray-200">
              <nav className="-mb-px flex space-x-8" aria-label="Tabs">
                {tabs.map((tab) => (
                  <SettingsTab
                    key={tab}
                    name={tab}
                    current={currentTab}
                    setCurrent={setCurrentTab}
                  />
                ))}
              </nav>
            </div>
          </div>
        </div>
      </div>

      <div className="mt-6">
        {currentTab === 'Account' && <AccountSettings />}
        {currentTab === 'API' && <APISettings />}
        {currentTab === 'Security & Privacy' && <SecuritySettings />}
      </div>
    </div>
  );
} 