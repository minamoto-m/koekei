# 🧾 声計 - 声でつける家計簿

声計は、音声入力とカレンダーUIで支出管理をサクッとできる家計簿アプリです。
バックエンドはGo、フロントエンドはNext.jsを使用しています。

---

## 📁 構成

koekei/
├── backend/   # Go APIサーバー（Gin + GORM）
├── frontend/  # Next.js フロントエンド

---

## 🚀 セットアップ

### 📦 フロントエンド（Next.js）

```bash
cd frontend
npm install
npm run dev
```

### 🧠 バックエンド（Go）

```bash
cd backend
go run main.go
```

---

## 🔧 技術スタック
- Go (Echo, GORM)
- Next.js (App Router)
- Whisper API（音声認識）
- PostgreSQL（開発用DB）