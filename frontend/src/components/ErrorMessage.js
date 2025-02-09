import React from 'react';

function ErrorMessage({ message }) {
  return (
    <div className="alert alert-danger" role="alert">
      {message}
    </div>
  );
}

export default ErrorMessage;