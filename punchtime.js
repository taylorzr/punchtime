window.hours = function () {
  return {
    day: 0,
    hours: [],
    core_hours: [],
    sortColumn: 'core_hours',
    sortDesc: true,
    columns: [
      { "code_name": "name", "pretty_name": "Name" },
      { "code_name": "hours", "pretty_name": "Hours" },
      { "code_name": "core_hours", "pretty_name": "Core Hours" },
      { "code_name": "not_core_hours", "pretty_name": "!Core Hours" },
      { "code_name": "first_punch", "pretty_name": "First Punch" },
      { "code_name": "last_punch", "pretty_name": "Last Punch" },
      { "code_name": "punch_count", "pretty_name": "Punch Count" }
    ],
    date() {
      d = new Date();
      d.setDate(d.getDate() + this.day)
      return d.toDateString()
    },
    getHours(day) {
      url = '/api/hours?day=' + day
      fetch(url)
        .then((res) => res.json())
        .then((data) => {
          console.log(data);
          this.hours = data.hours;
          this.sort();
        });
    },
    sort() {
      this.hours = this.hours.sort((a, b) => {
        column = this.sortColumn;
        if (column == 'name') {
          if (this.sortDesc) {
            return b[column] > a[column]
          } else {
            return a[column] > b[column]
          }
        } else {
          if (this.sortDesc) {
            return b[column] - a[column]
          } else {
            return a[column] - b[column]
          }
        }
      });
    },
    sortBy(column) {
      if (this.sortColumn == column) {
        this.sortDesc = !this.sortDesc;
      }
      this.sortColumn = column;
      this.sort();
    }
  };
}

window.punches = function () {
  return {
    punches: [],
    getPunches() {
      username = window.location.pathname.replace(/\/punches\//, '')
      fetch('/api/punches/' + username)
        .then((res) => res.json())
        .then((data) => {
          console.log(data);
          this.punches = data;
        });
    }
  };
}
