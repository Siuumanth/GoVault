# GoVault Frontend - Analysis & Fixes

## Issues Found & Solutions

### 1. **Date Sorting & Display** ✅
**Issue**: Files are not sorted by date, and dates are not displayed.

**Solution**: 
- Sort files by `created_at` or `updated_at` (whichever field exists in API response)
- Group files by date (Today, Yesterday, This Week, etc.) or show date header
- Display date above each file card

**Implementation**: Modify `Dashboard.jsx` to sort files after fetching, and update `FileCard.jsx` to display date.

---

### 2. **Update Name Option Not Working** ❌
**Issue**: The `rename` function exists in `files.js` but there's no UI component to trigger it.

**Root Cause**: 
- `FileCard` component doesn't have a rename button or modal
- `Dashboard` doesn't pass an `onRename` handler to `FileCard`

**Solution**:
- Add a rename button/icon in `FileCard` (e.g., pencil icon)
- Create a simple rename modal or inline input
- Connect it to `filesApi.rename(fileId, newName)`
- Refresh file list after rename

---

### 3. **Share Option Not Working** ❌
**Issue**: Share button exists but doesn't work.

**Root Cause**:
- `Dashboard.jsx` doesn't pass `onShare` handler to `FileCard`
- `ShareModal.jsx` references `ENDPOINTS.SHARING.LIST` which doesn't exist in `endpoints.js`
- Missing endpoint definitions for sharing API

**Solution**:
- Add `SHARING` endpoints to `endpoints.js`:
  ```js
  SHARING: {
    LIST: (fileId) => `/api/files/f/${fileId}/shares`,
    ADD: (fileId) => `/api/files/f/${fileId}/shares`,
    REMOVE: (fileId, userId) => `/api/files/f/${fileId}/shares/${userId}`,
  }
  ```
- Add share API functions to `files.js`
- Pass `onShare` handler from `Dashboard` to `FileCard`
- Fix `ShareModal` to use correct endpoints

---

### 4. **File Size Display** ✅
**Issue**: Always shows MB, even for small files (<1MB) or large files (>1GB).

**Current Code** (`FileCard.jsx` line 2):
```js
const formatSize = (bytes) => (bytes / (1024 * 1024)).toFixed(2) + " MB";
```

**Solution**: Update `formatSize` function:
```js
const formatSize = (bytes) => {
  if (bytes < 1024) return bytes + " B";
  if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(2) + " KB";
  if (bytes < 1024 * 1024 * 1024) return (bytes / (1024 * 1024)).toFixed(2) + " MB";
  return (bytes / (1024 * 1024 * 1024)).toFixed(2) + " GB";
};
```

---

### 5. **Make File Public / Delete Public Access** ❌
**Issue**: No UI or API integration for public access.

**Solution**:
- Add `PUBLIC` endpoints to `endpoints.js`:
  ```js
  PUBLIC: {
    CREATE: (fileId) => `/api/files/f/${fileId}/public`,
    DELETE: (fileId) => `/api/files/f/${fileId}/public`,
  }
  ```
- Add public API functions to `files.js`
- Add "Make Public" / "Remove Public" button in `FileCard` or a settings menu
- Show public URL when file is public

---

### 6. **Add/Delete Shortcut** ❌
**Issue**: Star button exists but functionality incomplete.

**Current**: `FileCard` has `onShortcut` prop but `Dashboard` doesn't pass it.

**Solution**:
- Add `SHORTCUT` endpoints:
  ```js
  SHORTCUT: {
    ADD: (fileId) => `/api/files/f/${fileId}/shortcut`,
    DELETE: (fileId) => `/api/files/f/${fileId}/shortcut`,
  }
  ```
- Add shortcut API functions to `files.js`
- Implement `handleAddShortcut` and `handleDeleteShortcut` in `Dashboard`
- Check if file is already a shortcut and toggle star icon state

---

### 7. **Website Title & Icon** ✅
**Issue**: Title is "client" and icon is default vite.svg.

**Solution**:
- Update `index.html`:
  - Change `<title>client</title>` to `<title>GoVault</title>`
  - Change icon href to placeholder URL (you'll add later)

---

## Implementation Plan

1. ✅ Fix file size formatting (quick fix)
2. ✅ Add date sorting and display
3. ✅ Fix share functionality (add endpoints, connect handlers)
4. ✅ Add rename functionality (create modal, connect API)
5. ✅ Add public access toggle
6. ✅ Fix shortcut add/delete
7. ✅ Update title and icon

All changes will preserve core upload logic and maintain existing functionality.
