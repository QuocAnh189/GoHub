const data = [
  {
    id: "1",
    name: "123",
    parentId: null,
    desc: "1",
    children: [],
  },
  {
    id: "2",
    name: "124",
    parentId: null,
    desc: "2",
    children: [],
  },
  {
    id: "3",
    name: "125",
    parentId: null,
    desc: "3",
    children: [
      {
        id: "4",
        name: "126",
        parentId: 3,
        desc: "4",
        children: [],
      },
    ],
  },
  {
    id: "5",
    name: "127",
    parentId: 1,
    desc: "5",
    children: [],
  },
  {
    id: "6",
    name: "128",
    parentId: 2,
    desc: "6",
    children: [],
  },
  {
    id: "7",
    name: "129",
    parentId: 3,
    desc: "7",
    children: [],
  },
];

const result = [
  {
    id: "1",
    name: "123",
    parentId: null,
    desc: "1",
    children: [
      {
        id: "5",
        name: "127",
        parentId: 1,
        desc: "5",
        children: [],
      },
    ],
  },
  {
    id: "2",
    name: "124",
    parentId: null,
    desc: "2",
    children: [
      {
        id: "6",
        name: "128",
        parentId: 2,
        desc: "6",
        children: [],
      },
    ],
  },
  {
    id: "3",
    name: "125",
    parentId: null,
    desc: "3",
    children: [
      {
        id: "4",
        name: "126",
        parentId: 3,
        desc: "4",
        children: [],
      },
      {
        id: "7",
        name: "129",
        parentId: 3,
        desc: "7",
        children: [],
      },
    ],
  },
];
