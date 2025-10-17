import { type Comment } from "../store/post";

export default function buildCommentTree(comments: Comment[]): any[] {
  const map = new Map<number, any>();
  const roots: any[] = [];

  comments.forEach((c) => map.set(c.id, { ...c, replies: [] }));
  console.log(comments);

  map.forEach((comment) => {
    if (comment.parent_id === 0) {
      roots.push(comment);
    } else {
      const parent = map.get(comment.parent_id);
      if (parent) {
        parent.replies.push(comment);
      }
    }
  });

  console.log(roots);

  return roots;
}
